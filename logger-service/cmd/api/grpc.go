package main

import (
	"context"
	"fmt"
	"github.com/bugrakocabay/logger-service/data"
	"github.com/bugrakocabay/logger-service/logs"
	"google.golang.org/grpc"
	"log"
	"net"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	res := &logs.LogResponse{Result: "logged"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPORT))
	if err != nil {
		log.Fatalf("failed to listen grpc server: %v", err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC Server started on port %s", grpcPORT)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to listen grpc server: %v", err)
	}
}
