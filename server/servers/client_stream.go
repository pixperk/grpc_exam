package servers

import (
	"io"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
)

func (s *ExamServiceServer) SubmitExamResults(stream exampb.ExamService_SubmitExamResultsServer) error {
	var (
		totalExams         int32
		totalMarksObtained int32
		totalPossibleMarks int32
		studentID          string
	)

	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				avg := float32(totalMarksObtained) / float32(totalPossibleMarks) * 100
				res := &exampb.SubmitExamResultResponse{
					StudentId:          studentID,
					TotalExams:         totalExams,
					TotalMarksObtained: totalMarksObtained,
					TotalPossibleMarks: totalPossibleMarks,
					AveragePercentage:  avg,
				}
				return stream.SendAndClose(res)
			}

			return err
		}

		if studentID == "" {
			studentID = req.StudentId
		}

		totalExams++
		totalMarksObtained += req.MarksObtained
		totalPossibleMarks += req.TotalMarks
	}
}
