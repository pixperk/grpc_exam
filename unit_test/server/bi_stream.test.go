package server_test

import (
    "context"
    "io"
    "testing"
    "github.com/pixperk/grpc_exam/proto/generated/exampb"
    "github.com/pixperk/grpc_exam/server/servers"
    "github.com/stretchr/testify/assert"
    "google.golang.org/grpc"
)

// mockBiStream implements exampb.ExamService_LiveExamQueryServer interface
type mockBiStream struct {
    grpc.ServerStream
    recvChan chan *exampb.GetExamResultRequest
    sendChan chan *exampb.GetExamResultResponse
    ctx context.Context
}

func newMockBiStream() *mockBiStream {
    return &mockBiStream{
        recvChan: make(chan *exampb.GetExamResultRequest, 10),
        sendChan: make(chan *exampb.GetExamResultResponse, 10),
        ctx: context.Background(),
    }
}

func (m *mockBiStream) Recv() (*exampb.GetExamResultRequest, error) {
    req, ok := <-m.recvChan
    if !ok {
        return nil, io.EOF
    }
    return req, nil
}

func (m *mockBiStream) Send(resp *exampb.GetExamResultResponse) error {
    m.sendChan <- resp
    return nil
}

func (m *mockBiStream) Context() context.Context {
    return m.ctx
}

func TestLiveExamQuery(t *testing.T) {
    server := servers.NewExamServiceServer()
    stream := newMockBiStream()

    // Running LiveExamQuery in a goroutine so that it will block reading Recv
    done := make(chan struct{})
    go func() {
        err := server.LiveExamQuery(stream)
        assert.NoError(t, err)
        close(done)
    }()

    //  To Send requests to recvChan to simulate client sending requests
    stream.recvChan <- &exampb.GetExamResultRequest{StudentId: "123", ExamId: "math101"}
    stream.recvChan <- &exampb.GetExamResultRequest{StudentId: "123", ExamId: "phy101"}

    // Closing the channel to simulate client closing the stream
    close(stream.recvChan)

    // Collect responses
    var responses []*exampb.GetExamResultResponse
    for i := 0; i < 2; i++ {
        resp := <-stream.sendChan
        responses = append(responses, resp)
    }

    <-done

    assert.Len(t, responses, 2)

    assert.Equal(t, "John Doe", responses[0].StudentName)
    assert.Equal(t, "Math 101", responses[0].Subject)
    assert.Equal(t, int32(95), responses[0].MarksObtained)

    assert.Equal(t, "John Doe", responses[1].StudentName)
    assert.Equal(t, "Physics 101", responses[1].Subject)
    assert.Equal(t, int32(81), responses[1].MarksObtained)
}