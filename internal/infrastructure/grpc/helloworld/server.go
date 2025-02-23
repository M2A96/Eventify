package helloworld

import (
	"context"

	proto "github.com/M2A96/Eventify.git/gen/go"
	"github.com/M2A96/Eventify.git/internal/core/ports/input"
)

type GRPCServer struct {
	service input.HelloworldServicePort
	proto.UnimplementedHelloServiceServer
}

func NewGRPCServer(service input.HelloworldServicePort) *GRPCServer {
	return &GRPCServer{service: service}
}

func (s *GRPCServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	response, err := s.service.Greet(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.HelloResponse{Message: response.Message}, nil
}
