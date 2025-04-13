package main

import (
	"log/slog"

	"github.com/pixperk/grpc_exam/client/clients"
	"github.com/pixperk/grpc_exam/proto/generated/exampb"
	"github.com/pixperk/grpc_exam/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	utils.InitLogger(true)
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to connect to server", "error", err)
	}
	defer conn.Close()

	client := exampb.NewExamServiceClient(conn)

	//clients.Unary(client)
	clients.Server_stream(client)

}
