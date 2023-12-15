//go:build !production

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"

	_ "github.com/traggo/server/docs"
)

func InitSwagger(app *fiber.App) {
	log.Debug().Msg("server.Swagger()")

	app.Get("/swagger/*", swagger.HandlerDefault)
}
