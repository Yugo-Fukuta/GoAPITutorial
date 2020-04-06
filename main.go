package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Yugo-Fukuta/GoAPITutorial/db"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})

	db.Init()
	r.Run()

	db.Close()
}
