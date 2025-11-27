package main

import (
	"fmt"
	"log"
	"os"
	"retroTV/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running!"})
	})

	// Load era JSON
	handlers.InitEra("./data/2000s.json")

	// Routes
	r.GET("/categories", handlers.CategoriesHandler)
	r.GET("/category", handlers.CategoryVideosHandler)
	r.GET("/random", handlers.RandomVideoHandler)

	return r
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return fmt.Sprintf(":%s", port)
}

func main() {
	log.Println("Starting server")
	r := setupRouter()
	port := getPort()
	log.Printf("Server running on http://localhost%s", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
