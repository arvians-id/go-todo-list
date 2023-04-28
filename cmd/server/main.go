package main

import (
	"fmt"
	"github.com/arvians-id/go-todo-list/cmd/config"
	"github.com/arvians-id/go-todo-list/internal"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	// Init Config
	configuration := config.New(".env")
	db, err := config.InitMySQL(configuration)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Init Server
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-API-KEY",
		AllowMethods:     "POST, DELETE, PUT, PATCH, GET",
		AllowCredentials: true,
	}))

	internal.InitRoutes(app, db)

	port := fmt.Sprintf(":%s", configuration.Get("PORT"))
	err = app.Listen(port)
	if err != nil {
		log.Fatalln("Cannot connect to server", err)
	}
}
