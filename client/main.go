package main

import (
	"context"
	"log"
	"time"

	pb "grpc-printer"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться: %v", err)
	}
	defer conn.Close()

	client := pb.NewPrinterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.PrintRequest{Text: "Привет, сервер!"}
	res, err := client.PrintMessage(ctx, req)
	if err != nil {
		log.Fatalf("Ошибка при вызове PrintMessage: %v", err)
	}

	log.Printf("Ответ сервера: %s", res.GetStatus())
}
