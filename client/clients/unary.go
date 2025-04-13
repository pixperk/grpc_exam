package clients

import (
	"context"
	"fmt"
	"time"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
)

func Unary(client exampb.ExamServiceClient) {

	fmt.Println("Enter student ID and exam ID (e.g., 123 math101):")
	var studentID, examID string
	fmt.Scanf("%s %s", &studentID, &examID)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.GetExamResult(ctx, &exampb.GetExamResultRequest{StudentId: studentID, ExamId: examID})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Student Name: %s\n", resp.StudentName)
	fmt.Printf("Subject: %s\n", resp.Subject)
	fmt.Printf("Marks Obtained: %d out of %d\n", resp.MarksObtained, resp.TotalMarks)
	fmt.Printf("Grade: %s\n", resp.Grade)
	fmt.Println("Unary RPC call completed successfully.")

}
