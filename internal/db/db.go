package db

import (
	"context"
	"database/sql"
	"time"
)

func New(addr string, maxConn int8, maxIdleConn int8, maxIdleTime string) (*sql.DB, error) {
	// Running set things up
	db, error := sql.Open("postgres", addr)
	if error != nil {
		return nil, error
	}

	db.SetMaxOpenConns(int(maxConn))
	db.SetConnMaxIdleTime(time.Duration(maxIdleConn))

	// Parsing from string to duration
	idleDuration, err := time.ParseDuration(maxIdleTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(idleDuration)

	// Tạo một context mới nhằm kiểm tra tính hiệu đến database có hoạt động hay không
	// Sử dụng context nhằm đảm bảo timeout cho việc connect và có giá trị trả về
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if error := db.PingContext(ctx); error != nil {
		return nil, error
	}

	return db, nil
}
