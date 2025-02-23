package input

import "github.com/M2A96/Eventify.git/internal/core/domain"

type HelloworldServicePort interface {
	Greet(name string) (*domain.HelloResponse, error)
}
