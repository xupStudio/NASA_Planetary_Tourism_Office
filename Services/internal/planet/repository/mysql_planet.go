package repository

import (
	"github.com/go-xorm/xorm"
	Planet "github.com/nasa/planetary_tourism/internal/planet"
	"github.com/nasa/planetary_tourism/models"
)

type PlanetRepo struct {
	orm *xorm.Engine
}

func NewPlanet(orm *xorm.Engine) Planet.Repository {
	return &PlanetRepo{
		orm: orm,
	}
}

func (m *PlanetRepo) GetPlanet(planetId string) (*models.Planet, error) {
	p := models.PlanetMap[planetId]
	return &p, nil

}

func (m *PlanetRepo) GetPlanetAttration(planetId string) ([]models.PlanetAttraction, error) {

	return nil, nil

}
