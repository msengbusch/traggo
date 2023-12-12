package config

import (
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

type EnvironmentConfig struct {
	values map[string]string
}

func (c *EnvironmentConfig) GetString(key string) string {
	value := c.values[key]
	if value == "" {
		// TODO better error
		log.Fatal().Msgf("Required config value %s is not set. Use environment variables or config files", key)
	}

	return value
}

func (c *EnvironmentConfig) GetOrDefaultString(key string, defaultValue string) string {
	value := c.values[key]
	if value == "" {
		return defaultValue
	}

	return value
}

func (c *EnvironmentConfig) GetInt(key string) int {
	value := c.values[key]
	if value == "" {
		// TODO better error
		log.Fatal().Msgf("Required config value %s is not set. Use environment variables or config files", key)
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal().Msgf("Config value %s is not an integer", key)
	}

	return intValue
}

func (c *EnvironmentConfig) GetOrDefaultInt(key string, defaultValue int) int {
	value := c.values[key]
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal().Msgf("Config value %s is not an integer", key)
	}

	return intValue
}

func (c *EnvironmentConfig) GetBool(key string) bool {
	value := c.values[key]
	if value == "" {
		// TODO better error
		log.Fatal().Msgf("Required config value %s is not set. Use environment variables or config files", key)
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatal().Msgf("Config value %s is not a boolean", key)
	}

	return boolValue
}

func (c *EnvironmentConfig) GetOrDefaultBool(key string, defaultValue bool) bool {
	value := c.values[key]
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatal().Msgf("Config value %s is not a boolean", key)
	}

	return boolValue
}

func Parse(lines []string) EnvironmentConfig {
	values := make(map[string]string, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		if line[0] == '#' {
			continue
		}

		var key string
		var value string

		before, after, found := strings.Cut(line, "=")
		if !found {
			key = before
			value = "true"
		} else {
			key = before
			value = after
		}

		if values[key] == "" {
			values[key] = value
		}
	}

	return EnvironmentConfig{values: values}
}
