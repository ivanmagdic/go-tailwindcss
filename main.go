package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/ivanmagdic/go-tailwindcss/router"
	"log"
)

func main() {
	engine := html.New("./resource/views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public/dist")

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}