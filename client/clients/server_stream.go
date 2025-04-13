package clients

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
)

func Server_stream(client exampb.ExamServiceClient) {

	req := &exampb.StreamExamResultsRequest{
		StudentId: "123",
		ExamIds:   []string{"math101", "phy101", "hist101"},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.StreamExamResults(ctx, req)
	if err != nil {
		log.Fatalf("error calling StreamExamResults: %v", err)
	}

	fmt.Println("Streaming exam results:")

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("All results received")
				break
			}
			log.Fatalf("error receiving exam result: %v", err)
		}

		fmt.Printf("- %s: %s (%d/%d), Grade: %s\n",
			res.StudentName, res.Subject, res.MarksObtained, res.TotalMarks, res.Grade)
	}
}
