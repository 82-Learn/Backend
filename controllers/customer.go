package controllers

import "github.com/gofiber/fiber/v2"

func Customers(c *fiber.Ctx) error {
	return c.SendString("new to this.")
}
