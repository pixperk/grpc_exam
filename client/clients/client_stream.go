package clients

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
)

func Client_stream(client exampb.ExamServiceClient) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Student ID:")
	studentID, _ := reader.ReadString('\n')
	studentID = strings.TrimSpace(studentID)

	stream, err := client.SubmitExamResults(context.Background())
	if err != nil {
		fmt.Println("Error creating stream:", err)
		return
	}

	for {
		fmt.Println("Enter subject (or leave blank to finish): ")
		subject, _ := reader.ReadString('\n')
		subject = strings.TrimSpace(subject)

		if subject == "" {
			break
		}

		fmt.Println("Enter Marks Obtained:")
		marksStr, _ := reader.ReadString('\n')
		marksStr = strings.TrimSpace(marksStr)
		marksObtained, _ := strconv.ParseInt(marksStr, 10, 32)

		fmt.Println("Enter Total Marks:")
		totalStr, _ := reader.ReadString('\n')
		totalStr = strings.TrimSpace(totalStr)
		total, _ := strconv.Atoi(totalStr)

		req := &exampb.SubmitExamResultRequest{
			StudentId:     studentID,
			Subject:       subject,
			MarksObtained: int32(marksObtained),
			TotalMarks:    int32(total),
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending request: %v", err)
		}

		fmt.Println("Request sent successfully")

	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}

	fmt.Println("\n📊 Summary:")
	fmt.Printf("Student ID: %s\n", res.StudentId)
	fmt.Printf("Total Exams: %d\n", res.TotalExams)
	fmt.Printf("Total Marks: %d/%d\n", res.TotalMarksObtained, res.TotalPossibleMarks)
	fmt.Printf("Average: %.2f%%\n", res.AveragePercentage)

}
