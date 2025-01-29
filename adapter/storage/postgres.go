package storage

import (
	"database/sql"
	"fmt"

	"github.com/praveen-shivalingaiah/go-url-shortner/internal/domain"
)

type PostgresURLRepository struct {
	db *sql.DB
}

func NewPostgresURLRepository(db *sql.DB) *PostgresURLRepository {
	return &PostgresURLRepository{db: db}
}

func (r *PostgresURLRepository) Save(u domain.URL) error {
	query := `INSERT INTO urls (short_id, original_url, created_at) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, u.ShortID, u.OriginalURL, u.CreatedAt)
	return err
}

func (r *PostgresURLRepository) FindOne(shortID string) (*domain.URL, error) {
	var url domain.URL
	query := `SELECT short_id, original_url, created_at FROM urls WHERE short_id = $1`
	row := r.db.QueryRow(query, shortID)
	if err := row.Scan(&url.ShortID, &url.OriginalURL, &url.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no URL found for short_id: %s", shortID)
		}
		return nil, err
	}
	return &url, nil
}
