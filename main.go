package main

import (
	_ "github.com/abiyyu03/siruta/config/timezone"

	"github.com/abiyyu03/siruta/config"
	"github.com/abiyyu03/siruta/delivery/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitDB()

	// seed := new(seeder.SeederStruct)
	// seed.RunSeeders()

	app := fiber.New()

	routes.HttpRoutes(app)

	app.Listen(":8080")
}
