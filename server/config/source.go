package config

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/traggo/server/util"
)

var (
	prefix        = "TRAGGO_"
	relativeFiles = []string{".env.development.local", ".env.development", ".env.local", ".env"}
	absoluteFiles = []string{"/etc/traggo/server_.ini"}
)

func getVars() []string {
	var vars = os.Environ()

	workingDirectory, err := util.GetWorkingDirectory()
	if err != nil {
		// TODO better error
		log.Fatal().Msgf("Cannot get working directory for config initialization: %s", err)
	}

	for _, file := range relativeFiles {
		content := readFile(workingDirectory + "/" + file)
		if content != nil {
			vars = append(vars, content...)
		}
	}

	for _, file := range absoluteFiles {
		content := readFile(file)
		if content != nil {
			vars = append(vars, content...)
		}
	}

	return vars
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug().Err(err).Msgf("Config file %s does not exist, ignoring it", path)
		} else {
			log.Warn().Err(err).Msgf("Cannot open config file %s, ignoring it", path)
		}
		return nil
	}

	lines, err := util.ReadLines(file)
	if err != nil {
		log.Warn().Err(err).Msgf("Cannot read config file %s, ignoring it", path)
		return nil
	}

	return lines
}
