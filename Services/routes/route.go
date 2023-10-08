package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	app := fiber.New()

	// cors setting
	//app.Use(cors.New(cors.Config{
	//	AllowOrigins: "http://localhost:9999",
	//	AllowHeaders: "Origin,Content-Type,Accept",
	//}))

	PlanetV1(app)
	return app
}
