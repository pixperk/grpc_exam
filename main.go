package main

import (
	"net"

	"log/slog"

	"github.com/pixperk/grpc_exam/utils"
	"google.golang.org/grpc"
)

func main() {
	utils.InitLogger(true)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("failed to listen", "error", err)
	}

	s := grpc.NewServer()

	//Register services

	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
	} else {
		slog.Info("server started successfully")
	}
}
