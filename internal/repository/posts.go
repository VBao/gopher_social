package repository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

// Model
type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserId    int64    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

// Có thể gồm tiền tố là database kết nối đến
// Trong trường hợp interface được share cho
// nhiều database khác nhau
type PostsRepository struct {
	db *sql.DB
}

// Repository thao tac voi database
func (s *PostsRepository) Create(ctx context.Context, post *Post) error {
	// Sử dụng cho Postgres
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserId,
		pq.Array(post.Tags),
	).Scan( // Map giá trị trả về từ RETURNING từ câu truy vấn
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
