package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "github.com/M2a96/Eventify/internal/core/services"
	proto "github.com/M2A96/Eventify.git/gen/go"
	helloworld "github.com/M2A96/Eventify.git/internal/core/services"
	grpcAdapter "github.com/M2A96/Eventify.git/internal/infrastructure/grpc/helloworld"
	httpAdapter "github.com/M2A96/Eventify.git/internal/infrastructure/http/helloworld"
	repoAdapter "github.com/M2A96/Eventify.git/internal/infrastructure/repository/helloworld"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

func main() {
	// Initialize dependencies
	repo := repoAdapter.NewMemoryRepository()
	service := helloworld.NewService(repo)

	// Setup gRPC server
	grpcServer := grpc.NewServer()
	grpcAdapter := grpcAdapter.NewGRPCServer(service)
	proto.RegisterHelloServiceServer(grpcServer, grpcAdapter)

	// Setup HTTP server
	e := echo.New()
	httpHandler := httpAdapter.NewHTTPHandler(service)
	e.GET("/hello/:name", httpHandler.GetHello)

	// Start servers
	go startGRPCServer(grpcServer)
	go startHTTPServer(e)

	// Graceful shutdown
	waitForShutdown(grpcServer, e)
}

func startGRPCServer(grpcServer *grpc.Server) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting gRPC server on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}

func startHTTPServer(e *echo.Echo) {
	fmt.Println("Starting HTTP server on :8080")
	if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
		e.Logger.Fatal("shutting down the server")
	}
}

func waitForShutdown(grpcServer *grpc.Server, e *echo.Echo) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("\nShutting down servers...")
	grpcServer.GracefulStop()
	if err := e.Close(); err != nil {
		e.Logger.Fatal(err)
	}
}
