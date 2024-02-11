package main

import (
	"cat-fact/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define your API endpoint
	r.GET("/cat-fact", handler.GetCatFact)

	// Run the server
	r.Run(":8080")
}
