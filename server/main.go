package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-printer"

	"google.golang.org/grpc"
)

type printerServer struct {
	pb.UnimplementedPrinterServiceServer
}

func (s *printerServer) PrintMessage(ctx context.Context, req *pb.PrintRequest) (*pb.PrintResponse, error) {
	fmt.Println("Получено сообщение:", req.GetText())
	return &pb.PrintResponse{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPrinterServiceServer(grpcServer, &printerServer{})

	log.Println("gRPC сервер запущен на порту 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
