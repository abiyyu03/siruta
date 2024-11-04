package main

import (
	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/delivery/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()
	app := fiber.New()
	//register routes
	routes.HttpRoutes(app)

	app.Listen(":8080")
}
