package storage

import (
	"fmt"
	"sync"

	"github.com/praveen-shivalingaiah/go-url-shortner/internal/domain"
)

type InMemoryURLRepository struct {
	data map[string]domain.URL
	mu   sync.RWMutex
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		data: make(map[string]domain.URL),
	}
}

func (r *InMemoryURLRepository) Save(url domain.URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[url.ShortID] = url
	return nil
}

func (r *InMemoryURLRepository) FindOne(shortID string) (*domain.URL, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if url, found := r.data[shortID]; found {
		return &url, nil
	}
	return nil, fmt.Errorf("URL not found: %s", shortID)
}
