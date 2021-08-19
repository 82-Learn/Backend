package main

import (
	"82learn.com/backend/database"
	"82learn.com/backend/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	defer database.DBConn.Close()

	routes.SetupRoutes(app)

	app.Listen(":3000")

}
