package main

import (
	"log"

	"github.com/rezakazemi928/social/internal/env"
)

func main(){
	conf := config{
		addr: env.GetString("ADDR", ":8080"),
	}
	app := &application{
		config: conf,
	}
	mo := app.mount()
	log.Fatal(app.run(mo))
}