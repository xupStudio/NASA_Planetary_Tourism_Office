package Planet

import "github.com/nasa/planetary_tourism/models"

type Repository interface {
	GetPlanet(string) (*models.Planet, error)
	GetPlanetAttration(string) ([]models.PlanetAttraction, error)
}
