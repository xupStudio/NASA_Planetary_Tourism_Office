package main

import (
	"log"

	"github.com/nasa/planetary_tourism/routes"
)

func main() {
	var err error

	// init env - orm
	// _ = api.InitOrm()

	// init env - route
	app := routes.Init()

	if err = app.Listen(":9999"); err != nil {
		log.Fatalf("launched error : %s ", err)
	}

}
