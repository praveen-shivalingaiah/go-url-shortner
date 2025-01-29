package app

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/praveen-shivalingaiah/go-url-shortner/internal/domain"
)

type ShortnerService struct {
	repo domain.URLRepository
}

func NewShortnerService(repo domain.URLRepository) *ShortnerService {
	return &ShortnerService{repo: repo}
}

func (s *ShortnerService) ShortenURL(original string) (string, error) {
	shortID := generateShortID()

	url := domain.URL{
		ShortID:     shortID,
		OriginalURL: original,
		CreatedAt:   time.Now(),
	}

	if err := url.Validate(); err != nil {
		return "", err
	}

	if err := s.repo.Save(url); err != nil {
		return "", err
	}

	return shortID, nil
}

func (s *ShortnerService) ResolveURL(shortID string) (string, error) {
	url, err := s.repo.FindOne(shortID)
	if err != nil {
		return "", err
	}

	if url == nil {
		return "", fmt.Errorf("URL not found")
	}

	return url.OriginalURL, nil
}

func generateShortID() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
