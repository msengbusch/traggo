package server

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() *fiber.App {
	log.Debug().Msg("Initializing server")

	app := fiber.New()
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
		Levels: []zerolog.Level{zerolog.ErrorLevel, zerolog.WarnLevel, zerolog.TraceLevel},
	}))

	return app
}

func Listen(app *fiber.App) {
	app.Listen(":4526")
}
