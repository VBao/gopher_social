package repository

import (
	"context"
	"database/sql"
)

// Có thể gồm tiền tố là database kết nối đến
// Trong trường hợp interface được share cho
// nhiều database khác nhau
type PostsRepository struct {
	db *sql.DB
}

func (s *PostsRepository) Create(ctx context.Context) error {
	return nil
}
