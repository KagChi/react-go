package cmd

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"react-go/routes"
)

func Init() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Prefork:     true,
	})

	app.Use(func(c *fiber.Ctx) error {
		err := c.Next()

		log.Info().
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", c.Response().StatusCode()).
			Str("remote_addr", c.IP()).
			Msg("Incoming request")

		return err
	})

	app.Static("/build", "./public/build")
	routes.RegisterWebRoutes(app)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal().Err(err).Msg("❌ Failed to start server")
	}
}
