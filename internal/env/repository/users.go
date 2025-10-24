package repository

import (
	"context"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func (s *UserRepository) Create(ctx context.Context) error {
	return nil
}
