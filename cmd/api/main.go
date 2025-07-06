package main

import "log"

func main(){
	conf := config{
		addr: ":8080",
	}
	app := &application{
		config: conf,
	}
	mo := app.mount()
	log.Fatal(app.run(mo))
}