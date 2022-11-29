package faucetRoutes

import (
	"github.com/gofiber/fiber/v2"
	faucetHandler "github.com/visveshdojima/faucet-backend/internals/handlers/Faucet"
)

func SetupFaucetRoutes(router fiber.Router) {
	faucet := router.Group("/faucet")
	// Create a Note
	faucet.Post("/", faucetHandler.CreateFaucetData)
	// Read all Notes
	// block.Get("/", blockHandler.GetNotes)
	// block.Delete("/", blockHandler.DeleteNotes)

	faucet.Get("/", faucetHandler.GetFaucetFromMongo)
	faucet.Post("/sendToken/:chain/:address", faucetHandler.SendToken)

}
