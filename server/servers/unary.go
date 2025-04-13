package servers

import (
	"context"
	"fmt"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
)

func (s *ExamServiceServer) GetExamResult(ctx context.Context, req *exampb.GetExamResultRequest) (*exampb.GetExamResultResponse, error) {
	key := fmt.Sprintf("%s_%s", req.StudentId, req.ExamId)
	if result, ok := s.examData[key]; ok {
		return result, nil
	} else {
		return nil, fmt.Errorf("exam result not found for student ID %s and exam ID %s", req.StudentId, req.ExamId)
	}
}
