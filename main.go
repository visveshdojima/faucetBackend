package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/visveshdojima/faucet-backend/database"
	"github.com/visveshdojima/faucet-backend/router"
)

func main() {
	fmt.Println("hrrl")

	app := fiber.New()
	database.ConnectMongoDB()
	router.SetupRoutes(app)

	app.Listen(":8080")
}
