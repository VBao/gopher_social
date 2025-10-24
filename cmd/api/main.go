package main

import (
	"log"

	"github.com/VBao/gopher_social/internal/env"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8888"),
	}

	app := &application{
		config: config,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
