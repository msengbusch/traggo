package main

import (
	"github.com/rs/zerolog/log"
	"github.com/traggo/server/database"
	"github.com/traggo/server/environment"
	"github.com/traggo/server/logger"
	"github.com/traggo/server/server"
)

// @title Traggo
// @version 1.0.0
// @host localhost:4526
// @BasePath /
func main() {
	logger.Init()
	printInformation()
	db := database.Init()
	app := server.Init()
	environment.Context.Init(app, db)
	server.Listen(app)
}

func printInformation() {
	log.Info().
		Str("env", string(environment.Mode)).
		Str("version", environment.BuildVersion).
		Str("commit", environment.BuildCommit).
		Str("date", environment.BuildDate).
		Msg("Traggo")
}
