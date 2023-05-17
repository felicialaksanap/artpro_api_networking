package main

import (
	"artpro_api_networking/db"
	"artpro_api_networking/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))

}
