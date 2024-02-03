package main

import (
	"fmt"
	"github.com/alijabbar034/database"
	"github.com/alijabbar034/routes"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	app := gin.Default()
	app.Static("/resources", "./resources")
	evnLoadError := godotenv.Load()
	if evnLoadError != nil {
		log.Fatal("env Load Error")

	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost/3000", "http://localhost:3001", "http://localhost:8000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		MaxAge:       time.Hour * 12,
	}))
	database.ConnectDB()
	app.Use()
	port := os.Getenv("PORT")
	api := app.Group("/api")

	routes.UserRouter(api)
	routes.ProductRouter(api)
	routes.ReviewRouter(api)
	routes.OrderRouter(api)

	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}
