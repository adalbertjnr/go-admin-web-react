package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/souzagmu/go-admin-web-react/handlers"
	"github.com/souzagmu/go-admin-web-react/middleware"
	"github.com/souzagmu/go-admin-web-react/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		load = handlers.NewStoreDB(store.ConnDB())
		app  = fiber.New()
	)

	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	app.Post("/api/register", load.Register)
	app.Post("/api/login", load.Login)
	app.Get("/api/user", middleware.JWTAuthenticate, load.User)
	app.Get("/api/logout", middleware.JWTAuthenticate, load.Logout)

	log.Fatal(app.Listen(":6000"))

}
