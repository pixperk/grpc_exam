package main

import (
	"log/slog"
	"os"

	"github.com/pixperk/grpc_exam/client/clients"
	"github.com/pixperk/grpc_exam/proto/generated/exampb"
	"github.com/pixperk/grpc_exam/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	utils.InitLogger(true)

	if len(os.Args) < 2 {
		slog.Error("Usage: go run client/main.go [unary|server|client|bi]")
		return
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("Failed to connect to server", "error", err)
		return
	}
	defer conn.Close()

	client := exampb.NewExamServiceClient(conn)

	switch os.Args[1] {
	case "unary":
		clients.Unary(client)
	case "server":
		clients.Server_stream(client)
	case "client":
		clients.Client_stream(client)
	case "bi":
		clients.BiDirectional(client)
	default:
		slog.Error("Unknown command. Use one of: unary, server, client, bi")
	}
}
