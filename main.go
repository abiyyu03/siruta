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
	// seed.Run()

	app := fiber.New()
	//register routes
	routes.HttpRoutes(app)

	app.Listen(":8080")
}
