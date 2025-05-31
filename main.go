package main

import (
	"data-dashboard/cliArgs"
	"data-dashboard/dbInteractions"
	"data-dashboard/handlers"

	_ "github.com/mattn/go-sqlite3"
)


func main () {
	cliargs.InitArgs()
	dbInteractions.InitDB()


	handlers.InitHandlers()
}
