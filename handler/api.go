package handler

import "github.com/gofiber/fiber/v2"

func Response(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success",
		"message": "Foo Bar",
		"data": nil,
	})
}