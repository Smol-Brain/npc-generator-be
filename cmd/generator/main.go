package main

import (
	"npc-generator-be/app"
)

func main() {
	// Load config variables before initializing other components
	conf := app.InitializeConfig()

	// Set up database, server
	db := app.InitializeDB(conf)
	app.InitializeServer(conf, db)
}
