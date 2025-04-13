package servers

import "github.com/pixperk/grpc_exam/proto/generated/exampb"

type ExamServiceServer struct {
	exampb.UnimplementedExamServiceServer
	examData map[string]*exampb.GetExamResultResponse
}

func NewExamServiceServer() *ExamServiceServer {
	data := map[string]*exampb.GetExamResultResponse{
		"123_math101": {
			StudentName:   "John Doe",
			Subject:       "Math 101",
			MarksObtained: 95,
			TotalMarks:    100,
			Grade:         "A+",
		},
		"123_phy101": {
			StudentName:   "John Doe",
			Subject:       "Physics 101",
			MarksObtained: 81,
			TotalMarks:    100,
			Grade:         "A",
		},
		"123_hist101": {
			StudentName:   "John Doe",
			Subject:       "History 101",
			MarksObtained: 76,
			TotalMarks:    100,
			Grade:         "B",
		},
		"456_phy101": {
			StudentName:   "Jane Smith",
			Subject:       "Physics 101",
			MarksObtained: 88,
			TotalMarks:    100,
			Grade:         "A",
		},
		"789_chem101": {
			StudentName:   "Alice Johnson",
			Subject:       "Chemistry 101",
			MarksObtained: 92,
			TotalMarks:    100,
			Grade:         "A+",
		},
		"101_bio101": {
			StudentName:   "Bob Brown",
			Subject:       "Biology 101",
			MarksObtained: 85,
			TotalMarks:    100,
			Grade:         "A",
		},
		"102_hist101": {
			StudentName:   "Charlie Davis",
			Subject:       "History 101",
			MarksObtained: 90,
			TotalMarks:    100,
			Grade:         "A+",
		},
	}

	return &ExamServiceServer{
		examData: data,
	}
}
