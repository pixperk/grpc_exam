package servers

import (
	"fmt"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
)

func (s *ExamServiceServer) StreamExamResults(req *exampb.StreamExamResultsRequest, stream exampb.ExamService_StreamExamResultsServer) error {
	studentId := req.StudentId
	examIDs := req.ExamIds

	found := false

	for _, examID := range examIDs {
		key := fmt.Sprintf("%s_%s", studentId, examID)
		if result, ok := s.examData[key]; ok {
			stream.Send(result)
			found = true
		}

		if !found {
			return fmt.Errorf("exam results not found for student ID %s and exam IDs %v", studentId, examIDs)
		}

	}

	return nil

}
