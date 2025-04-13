package servers

import (
	"fmt"
	"io"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
)

func (s *ExamServiceServer) LiveExamQuery(stream exampb.ExamService_LiveExamQueryServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		key := fmt.Sprintf("%s_%s", req.StudentId, req.ExamId)
		result, ok := s.examData[key]
		if !ok {
			err := stream.Send(&exampb.GetExamResultResponse{
				StudentName:   "N/A",
				Subject:       req.ExamId,
				MarksObtained: 0,
				TotalMarks:    0,
				Grade:         "Not Found",
			})
			if err != nil {
				return err
			}
			continue
		}

		if err := stream.Send(result); err != nil {
			return err
		}
	}
}
