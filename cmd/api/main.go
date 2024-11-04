package main

import "log"

func main() {
	svconf := serverConfig{
		addr: ":8080",
	}

	app := &application{
		config: svconf,
	}
	mux := app.mount()
	log.Fatal(app.start(mux))

}
