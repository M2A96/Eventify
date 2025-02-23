package helloworld

import (
	"fmt"

	"github.com/M2A96/Eventify.git/internal/core/domain"
	"github.com/M2A96/Eventify.git/internal/core/ports/input"
	"github.com/M2A96/Eventify.git/internal/core/ports/output"
)

type HelloworldService struct {
	repo output.HelloworldRepository
}

func NewService(repo output.HelloworldRepository) input.HelloworldServicePort {
	return &HelloworldService{repo: repo}
}

func (s *HelloworldService) Greet(name string) (*domain.HelloResponse, error) {
	existing, err := s.repo.GetGreeting(name)
	if err == nil {
		return &domain.HelloResponse{Message: existing.Message}, nil
	}

	newGreeting := &domain.Hello{
		Name:    name,
		Message: fmt.Sprintf("Hello, %s!", name),
	}

	if err := s.repo.SaveGreeting(newGreeting); err != nil {
		return nil, fmt.Errorf("failed to save greeting: %w", err)
	}

	return &domain.HelloResponse{Message: newGreeting.Message}, nil
}
