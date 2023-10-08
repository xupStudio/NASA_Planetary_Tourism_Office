package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nasa/planetary_tourism/internal/planet/repository"
	"github.com/nasa/planetary_tourism/internal/planet/service"
)

func GetPlanet(ctx *fiber.Ctx) error {
	var err error
	idStr := ctx.Params("planet_id")

	merRepo := repository.NewPlanet(env.orm)
	merSrv := service.NewPlanet(merRepo)
	resp, err := merSrv.GetPlanet(idStr)
	if err != nil {
		return err
	}

	_ = ctx.JSON(resp)
	return err
}

func GetPlanetAttration(ctx *fiber.Ctx) error {
	// var err error
	// idStr := ctx.Params("planet_id")

	// merRepo := repository.NewPlanet(env.orm)
	// merSrv := service.NewPlanet(merRepo)
	// mers, err := merSrv.FindPlanet(merCondi)
	// if err != nil {
	// 	return err
	// }

	// _ = ctx.JSON(mers)

	// return err
	return nil
}
