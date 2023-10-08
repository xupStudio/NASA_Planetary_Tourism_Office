package models

type Planet struct {
	Type                   string  `json:"type"`
	Size                   int     `json:"size"`
	RelatedSizeOfEarth     float64 `json:"relatedSizeOfEarth"`
	Mass                   string  `json:"mass"`
	RealtedMassOfEarth     float64 `json:"relatedMassOfEarth"`
	LengthOfAYear          float64 `json:"lengthOfAYear"`
	NumberOfMoons          int     `json:"numerOfMoons"`
	AverageDistanceFromSun float64 `json:"averageDistanceFromSun"`
	Temperature            string  `json:"temerature"`
	Formation              string  `json:"formation"`
	OribitAndRotation      string  `json:"orbitAndRotation"`
	Seasons                string  `json:"seasons"`
	Surface                string  `json:"surface"`
	Exosphere              string  `json:"exosphere"`
	Else                   string  `json:"else"`
}

type PlanetAttraction struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}
