package main

import (
	"net"

	"log/slog"

	"github.com/pixperk/grpc_exam/proto/generated/exampb"
	"github.com/pixperk/grpc_exam/server/servers"

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
	exampb.RegisterExamServiceServer(s, servers.NewExamServiceServer())

	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
	}
	slog.Info("server started successfully")

}
