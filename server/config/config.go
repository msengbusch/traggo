package config

import (
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port               int
	LogLevel           string
	DefaultUserName    string
	DefaultUserPass    string
	DatabaseDialect    string
	DatabaseConnection string
}

func Init() Config {
	log.Debug().Msg("Loading config")

	config := Config{}

	return config
}
