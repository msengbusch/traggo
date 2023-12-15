package server

import (
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() *fiber.App {
	log.Debug().Msg("server.Init()")

	app := fiber.New()

	logger(app)
	static(app)

	return app
}

func Listen(app *fiber.App) {
	app.Listen(":4526")
}

func logger(app *fiber.App) {
	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
		Levels: []zerolog.Level{zerolog.ErrorLevel, zerolog.WarnLevel, zerolog.TraceLevel},
	}))
}
