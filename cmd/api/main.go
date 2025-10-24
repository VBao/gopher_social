package main

import (
	"log"

	"github.com/VBao/gopher_social/internal/env"
	"github.com/VBao/gopher_social/internal/env/repository"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8888"),
	}

	repository := repository.NewRepository(nil)

	app := &application{
		config:     config,
		repository: repository,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
