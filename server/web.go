package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/traggo/server/web"
	"net/http"
)
import "github.com/gofiber/fiber/v2/middleware/filesystem"

func static(app *fiber.App) {
	log.Debug().Msg("server.Static()")
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(web.Files),
		PathPrefix: "dist",
		Browse:     true,
	}))
}
