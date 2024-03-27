package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port if not specified
	}

	r := gin.Default()

	// Middleware similar to morgan('dev') in Express
	r.Use(gin.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Add ID: %s, Home Page", port)
	})

	r.GET("/app1", func(c *gin.Context) {
		c.String(http.StatusOK, "Add ID: %s, App1 Page", port)
	})

	r.GET("/app2", func(c *gin.Context) {
		c.String(http.StatusOK, "Add ID: %s, App2 Page", port)
	})

	r.Run(":" + port) // Listen and serve on the specified port
}
