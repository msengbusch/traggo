package logger

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

type DatabaseLogger struct {
}

var _ logger.Interface = &DatabaseLogger{}

func (l *DatabaseLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

func (l *DatabaseLogger) Trace(c context.Context, begin time.Time, fc func() (sql string, rows int64), err error) {
	if log.Logger.GetLevel() > zerolog.TraceLevel {
		return
	}

	var event *zerolog.Event
	if err != nil {
		event = log.Debug()
	} else {
		event = log.Trace()
	}

	event.Ctx(c)
	event.Dur("elapsed", time.Since(begin))

	sql, rows := fc()
	if sql != "" {
		event.Str("sql", sql)
	}

	if rows >= 0 {
		event.Int64("rows", rows)
	}

	event.Send()
}

func (l *DatabaseLogger) Info(c context.Context, message string, args ...interface{}) {
	log.Info().Ctx(c).Msgf(message, args...)
}

func (l *DatabaseLogger) Warn(c context.Context, message string, args ...interface{}) {
	log.Warn().Ctx(c).Msgf(message, args...)
}

func (l *DatabaseLogger) Error(c context.Context, message string, args ...interface{}) {
	log.Error().Ctx(c).Msgf(message, args...)
}
