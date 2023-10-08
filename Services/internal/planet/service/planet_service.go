package service

import (
	Planet "github.com/nasa/planetary_tourism/internal/planet"
	"github.com/nasa/planetary_tourism/models"
	"github.com/nasa/planetary_tourism/pkg/errs"
)

type PlanetSvc struct {
	mRepo Planet.Repository
}

func NewPlanet(mRepo Planet.Repository) Planet.Service {
	return &PlanetSvc{
		mRepo: mRepo,
	}
}

func (m *PlanetSvc) GetPlanet(planetId string) (*models.Planet, error) {
	planet, err := m.mRepo.GetPlanet(planetId)
	if err != nil {
		return nil, errs.NewAppError(500, errs.DBGetErr, "", err)
	}

	return planet, nil

}

func (m *PlanetSvc) GetPlanetAttration(planetId string) ([]models.PlanetAttraction, error) {
	planetAttr, err := m.mRepo.GetPlanetAttration(planetId)
	if err != nil {
		return nil, errs.NewAppError(500, errs.DBGetErr, "", err)
	}

	return planetAttr, nil

}
