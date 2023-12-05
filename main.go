package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/souzagmu/go-admin-web-react/handlers"
	"github.com/souzagmu/go-admin-web-react/store"
)

func main() {

	var (
		load = handlers.NewStoreDB(store.ConnDB())
		app  = fiber.New()
	)

	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	app.Post("/api/register", load.Register)
	app.Post("/api/login", load.Login)

	log.Fatal(app.Listen(":6000"))

}
