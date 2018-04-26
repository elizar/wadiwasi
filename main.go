package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/", home)

	// Default port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// Run app
	r.Run(fmt.Sprintf(":%s", PORT))
}

func home(c *gin.Context) {
	c.String(200, "Hello, Universe!")
}
