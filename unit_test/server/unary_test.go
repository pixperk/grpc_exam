package server_test

import (
	"context"
	"testing"
	"github.com/pixperk/grpc_exam/proto/generated/exampb"
	"github.com/pixperk/grpc_exam/server/servers"
	"github.com/stretchr/testify/assert"
)

func TestGetExamResult(t *testing.T) {
	s := servers.NewExamServiceServer()

	tests := []struct {
		name string
		request  *exampb.GetExamResultRequest
		expected *exampb.GetExamResultResponse
	}{
		{
			name: "valid student and exam",
			request: &exampb.GetExamResultRequest{
				StudentId: "123",
				ExamId: "math101",
			},
			expected: &exampb.GetExamResultResponse{
				StudentName: "John Doe",
				Subject: "Math 101",
				MarksObtained: 95,
				TotalMarks: 100,
				Grade: "A+",
			},
		},
		{
			name: "unknown student/exam",
			request: &exampb.GetExamResultRequest{
				StudentId: "000",
				ExamId: "unknown",
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := s.GetExamResult(context.Background(), tt.request)

			if tt.expected == nil {
				assert.Error(t, err)
				assert.Nil(t, resp)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, resp)
			}
		})
	}
}