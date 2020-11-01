package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ivanmagdic/go-tailwindcss/handler"
)
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())
	api.Get("/", handler.Response)


	app.Get("/", handler.WelcomePage)
}