package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"

	faucetRoutes "github.com/visveshdojima/faucet-backend/internals/routes"
)

func SetupRoutes(app *fiber.App) {

	// Group api calls with param '/api'
	api := app.Group("/api", logger.New())

	faucetRoutes.SetupFaucetRoutes(api)

	// Setup note routes, can use same syntax to add routes for more models
}
