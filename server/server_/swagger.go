//go:build !production

package server_

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"

	_ "github.com/traggo/server/docs_"
)

func InitSwagger(app *fiber.App) {
	log.Debug().Msg("server_.Swagger()")

	app.Get("/swagger/*", swagger.HandlerDefault)
}
