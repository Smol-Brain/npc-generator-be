package main

import (
	"npc-generator-be/app"
	"npc-generator-be/config"
	"npc-generator-be/global"
)

func main() {
	// Load config variables before initializing other components
	// pwd, _ := os.Getwd()
	// fmt.Println(pwd)
	global.Conf, _ = config.InitializeConfig()

	// Set up database, server
	app.InitializeDB()
	app.InitializeRoutes()
}
