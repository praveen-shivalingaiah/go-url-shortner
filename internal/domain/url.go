package domain

import (
	"errors"
	"time"
)

var (
	UrlError = errors.New("Invalid URL")
)

type URL struct {
	ShortID     string
	OriginalURL string
	CreatedAt   time.Time
}

func (url *URL) Validate() error {
	if url.OriginalURL == "" {
		return UrlError
	}

	return nil
}

type URLRepository interface {
	Save(url URL) error
	FindOne(shortID string) (*URL, error)
}
