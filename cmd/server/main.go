package main

import (
	"log"

	"github.com/Guglahai/tasks-service/internal/configs"
	"github.com/Guglahai/tasks-service/internal/database"
	"github.com/Guglahai/tasks-service/internal/task"
	transportgrpc "github.com/Guglahai/tasks-service/internal/transport/grpc"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// Load configuration
	config := configs.New()

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	repo := task.NewRepository(db)
	svc := task.NewService(repo)

	userClient, conn, err := transportgrpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to create user client: %v", err)
	}
	defer conn.Close()

	if err := transportgrpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}
