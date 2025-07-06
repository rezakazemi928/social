package main

import "log"

func main(){
	conf := config{
		addr: ":8080",
	}
	app := &application{
		config: conf,
	}
	log.Fatal(app.run())
}