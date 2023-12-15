package main

import "github.com/rs/zerolog/log"
import "github.com/traggo/server/logger"

func main() {
	logger.Init()

	log.Info().Msg("Traggo CLI")
}
