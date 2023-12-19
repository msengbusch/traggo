package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"

	"github.com/traggo/server/logger"
)

func Init() *gorm.DB {
	log.Debug().Msg("database.Init()")

	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{
		Logger: &logger.DatabaseLogger{},
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect database")
	}

	return db
}
