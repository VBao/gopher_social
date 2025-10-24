package main

import "log"

func main() {
	config := config{
		addr: ":8098",
	}

	app := &application{
		config: config,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
