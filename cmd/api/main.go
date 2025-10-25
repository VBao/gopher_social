package main

import (
	"log"

	"github.com/VBao/gopher_social/internal/db"
	"github.com/VBao/gopher_social/internal/env"
	"github.com/VBao/gopher_social/internal/repository"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8888"),
		db: dbConfig{
			addr:        env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxConn:     int8(env.GetInt("DB_MAX_CONN", 30)),
			maxIdleConn: int8(env.GetInt("DB_MAX_IDLE_CONN", 10)),
			maxIdleTime: env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		config.db.addr,
		config.db.maxConn,
		config.db.maxIdleConn,
		config.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("Database connected completed!")

	repository := repository.NewRepository(db)

	app := &application{
		config:     config,
		repository: repository,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
