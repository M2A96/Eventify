package output

import "github.com/M2A96/Eventify.git/internal/core/domain"

type HelloworldRepository interface {
	SaveGreeting(greeting *domain.Hello) error
	GetGreeting(name string) (*domain.Hello, error)
}
