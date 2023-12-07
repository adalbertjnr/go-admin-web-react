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
	app.Get("/api/logout", middleware.JWTAuthenticate, load.Logout)

	app.Get("/api/users", middleware.JWTAuthenticate, load.GetAllUsers)
	app.Post("/api/user", middleware.JWTAuthenticate, load.CreateUser)
	app.Get("/api/user/:id", middleware.JWTAuthenticate, load.GetUser)
	app.Put("/api/user/:id", middleware.JWTAuthenticate, load.UpdateUser)
	app.Delete("/api/user/:id", middleware.JWTAuthenticate, load.DeleteUser)

	app.Get("/api/roles", middleware.JWTAuthenticate, load.GetAllRoles)
	app.Post("/api/role", middleware.JWTAuthenticate, load.CreateRole)
	app.Get("/api/role/:id", middleware.JWTAuthenticate, load.GetRole)
	app.Put("/api/role/:id", middleware.JWTAuthenticate, load.UpdateRole)
	app.Delete("/api/role/:id", middleware.JWTAuthenticate, load.DeleteRole)

	log.Fatal(app.Listen(":6000"))

}
