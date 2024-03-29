package main

import (
	"fmt"
	"github.com/Evgeny-Tokarev/go-serv/controller"
	"github.com/Evgeny-Tokarev/go-serv/database"
	"github.com/Evgeny-Tokarev/go-serv/middleware"
	"github.com/Evgeny-Tokarev/go-serv/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func serveApplication() {
	router := gin.Default()
	router.Use(cors.Default())

	publicRoutes := router.Group("/api/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.GET("/user", controller.GetCurrentUser)
	protectedRoutes.POST("/user/:id", controller.UpdateUser)
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
