package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nasa/planetary_tourism/api"
)

func PlanetV1(app *fiber.App) {
	v1 := app.Group("/v1")

	// GetPlanet
	v1.Get("/planet_info/:planet_id", func(ctx *fiber.Ctx) error {
		return api.GetPlanet(ctx)
	})

	// FindPlanet
	v1.Get("/planet_attration/:planet_id", func(ctx *fiber.Ctx) error {
		return api.GetPlanetAttration(ctx)
	})

}
