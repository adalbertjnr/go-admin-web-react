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

	api := app.Group("/api")

	api.Use(cors.New(cors.Config{AllowCredentials: true}))

	api.Post("/register", load.Register)
	api.Post("/login", load.Login)

	api.Put("/users/info", middleware.JWTAuthenticate, load.UpdateInfo)
	api.Put("/users/password", middleware.JWTAuthenticate, load.UpdatePassword)

	api.Get("/logout", middleware.JWTAuthenticate, load.Logout)

	api.Get("/users", middleware.JWTAuthenticate, load.GetAllUsers)
	api.Post("/user", middleware.JWTAuthenticate, load.CreateUser)
	api.Get("/user/:id", middleware.JWTAuthenticate, load.GetUser)
	api.Put("/user/:id", middleware.JWTAuthenticate, load.UpdateUser)
	api.Delete("/user/:id", middleware.JWTAuthenticate, load.DeleteUser)

	api.Post("/product", middleware.JWTAuthenticate, load.CreateProduct)
	api.Get("/products", middleware.JWTAuthenticate, load.GetAllProducts)
	api.Get("/product/:id", middleware.JWTAuthenticate, load.GetProduct)
	api.Put("/product/:id", middleware.JWTAuthenticate, load.UpdateProduct)
	api.Delete("/product/:id", middleware.JWTAuthenticate, load.DeleteProduct)

	api.Get("/roles", middleware.JWTAuthenticate, load.GetAllRoles)
	api.Post("/role", middleware.JWTAuthenticate, load.CreateRole)
	api.Get("/role/:id", middleware.JWTAuthenticate, load.GetRole)
	api.Put("/role/:id", middleware.JWTAuthenticate, load.UpdateRole)
	api.Delete("/role/:id", middleware.JWTAuthenticate, load.DeleteRole)

	api.Get("/permissions", middleware.JWTAuthenticate, load.GetAllPermissions)

	api.Post("/upload", middleware.JWTAuthenticate, load.Upload)

	log.Fatal(app.Listen(":6000"))

}
