package handler

import "github.com/gofiber/fiber/v2"

func WelcomePage(c *fiber.Ctx) error {
	/*return c.Render("welcome", fiber.Map{
		"Title": "Hello, World!",
	})*/

	return c.Render("welcome", fiber.Map{
	})
}
