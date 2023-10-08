package Planet

import "github.com/nasa/planetary_tourism/models"

type Service interface {
	GetPlanet(string) (*models.Planet, error)
	GetPlanetAttration(string) ([]models.PlanetAttraction, error)
}
