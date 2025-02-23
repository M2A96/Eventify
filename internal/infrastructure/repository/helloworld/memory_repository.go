package helloworld

import (
	"fmt"
	"sync"

	"github.com/M2A96/Eventify.git/internal/core/domain"
)

type MemoryRepository struct {
	mu    sync.RWMutex
	store map[string]*domain.Hello
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		store: make(map[string]*domain.Hello),
	}
}

func (r *MemoryRepository) SaveGreeting(greeting *domain.Hello) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.store[greeting.Name] = greeting
	return nil
}

func (r *MemoryRepository) GetGreeting(name string) (*domain.Hello, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if greeting, exists := r.store[name]; exists {
		return greeting, nil
	}
	return nil, fmt.Errorf("greeting not found")
}
