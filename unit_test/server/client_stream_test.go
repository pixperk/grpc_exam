package server_test

import (
	"testing"
	"io"
	"math"
	"github.com/pixperk/grpc_exam/proto/generated/exampb"
	"github.com/pixperk/grpc_exam/server/servers"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type fakeClientStream struct {
	grpc.ServerStream
	requests []*exampb.SubmitExamResultRequest
	index int
	response *exampb.SubmitExamResultResponse
}

func (f *fakeClientStream) Recv() (*exampb.SubmitExamResultRequest, error) {
	if f.index >= len(f.requests) {
		return nil, io.EOF
	}
	req := f.requests[f.index]
	f.index++
	return req, nil
}

func (f *fakeClientStream) SendAndClose(resp *exampb.SubmitExamResultResponse) error {
	f.response = resp
	return nil
}

func TestSubmitExamResults(t *testing.T) {
	server := servers.NewExamServiceServer()

	t.Run("submit_multiple_results", func(t *testing.T) {
		stream := &fakeClientStream{
			requests: []*exampb.SubmitExamResultRequest{
				{StudentId: "321", Subject: "Math", MarksObtained: 80, TotalMarks: 100},
				{StudentId: "321", Subject: "Physics", MarksObtained: 90, TotalMarks: 100},
				{StudentId: "321", Subject: "Biology", MarksObtained: 85, TotalMarks: 100},
			},
		}

		err := server.SubmitExamResults(stream)
		assert.NoError(t, err)
		assert.NotNil(t, stream.response)
		assert.Equal(t, "321", stream.response.StudentId)
		assert.Equal(t, int32(3), stream.response.TotalExams)
		assert.Equal(t, int32(255), stream.response.TotalMarksObtained)
		assert.Equal(t, int32(300), stream.response.TotalPossibleMarks)
		assert.InDelta(t, 85.0, stream.response.AveragePercentage, 0.01)
	})

	t.Run("no_results_submitted", func(t *testing.T) {
		stream := &fakeClientStream{requests: []*exampb.SubmitExamResultRequest{}}
		err := server.SubmitExamResults(stream)
		assert.NoError(t, err)
		assert.NotNil(t, stream.response)
		assert.Equal(t, int32(0), stream.response.TotalExams)
		assert.True(t, math.IsNaN(float64(stream.response.AveragePercentage)), "Expected NaN when no exam results are submitted")
	})
}
