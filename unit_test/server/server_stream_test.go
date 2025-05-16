package server_test

import (
    "testing"
    "github.com/pixperk/grpc_exam/proto/generated/exampb"
    "github.com/pixperk/grpc_exam/server/servers"
    "github.com/stretchr/testify/assert"
    "google.golang.org/grpc"
)

type fakeStream struct {
    grpc.ServerStream
    Responses []*exampb.GetExamResultResponse
}

func (fs *fakeStream) Send(resp *exampb.GetExamResultResponse) error {
    fs.Responses = append(fs.Responses, resp)
    return nil
}

func TestStreamExamResults(t *testing.T) {
    server := servers.NewExamServiceServer()

    t.Run("valid_student_multiple_exams", func(t *testing.T) {
        req := &exampb.StreamExamResultsRequest{
            StudentId: "123",
            ExamIds: []string{"math101", "phy101", "hist101"},
        }

        stream := &fakeStream{}

        err := server.StreamExamResults(req, stream)
        assert.NoError(t, err)
        assert.Len(t, stream.Responses, 3)

        assert.Equal(t, "Math 101", stream.Responses[0].Subject)
        assert.Equal(t, "Physics 101", stream.Responses[1].Subject)
        assert.Equal(t, "History 101", stream.Responses[2].Subject)
    })

    t.Run("some_exams_missing", func(t *testing.T) {
        req := &exampb.StreamExamResultsRequest{
            StudentId: "123",
            ExamIds: []string{"math101", "invalid101", "phy101"},
        }

        stream := &fakeStream{}

     err := server.StreamExamResults(req, stream)
     assert.NoError(t, err)          
    assert.Len(t, stream.Responses, 2)  
    assert.Equal(t, "Math 101", stream.Responses[0].Subject)
    assert.Equal(t, "Physics 101", stream.Responses[1].Subject)
    })

    t.Run("unknown_student", func(t *testing.T) {
        req := &exampb.StreamExamResultsRequest{
            StudentId: "999",
            ExamIds: []string{"math101", "phy101"},
        }

        stream := &fakeStream{}

        err := server.StreamExamResults(req, stream)
        assert.Error(t, err)
        assert.Nil(t, stream.Responses)
    })
}