package main

import (
	"github.com/Yugo-Fukuta/GoAPITutorial/db"
	"github.com/Yugo-Fukuta/GoAPITutorial/server"
)

func main() {
	db.Init()

	server.Init()

	db.Close()
}
