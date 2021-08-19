package controllers

import (
	"82learn.com/backend/database"
	"82learn.com/backend/models"
	"github.com/gofiber/fiber/v2"
)

func Customers(c *fiber.Ctx) error {
	return c.SendString("Try this again.")
}

func GetCustomers(c *fiber.Ctx) error {
	db := database.DBConn
	var customers []models.Customer
	db.Find(&customers)
	c.JSON(customers)
	return c.SendString("All customers")
}
