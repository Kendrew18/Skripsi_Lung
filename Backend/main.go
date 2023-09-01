package main

import (
	"kendrew/db"
	"kendrew/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":38600"))
}
