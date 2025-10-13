package logging

import (
	"os"
	"sync"

	"github.com/AmirHosein-Kahrani/Car-Center-Web/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var once sync.Once
var zeroSinLogger *zerolog.Logger

var zerologLevelMapping = map[string]zerolog.Level{

	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func newZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel() zerolog.Level {

	level, exists := zerologLevelMapping[l.cfg.Logger.Level]
	if !exists {
		return zerolog.DebugLevel
	}
	return level
}

func (l *zeroLogger) Init() {

	once.Do(func() {

		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

		file, err := os.OpenFile(l.cfg.Logger.Filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

		if err != nil {
			panic("zero Log Error: could not open file")
		}

		var logger = zerolog.New(file).
			With().
			Timestamp().
			Str("AppName", "MyApp").
			Str("LoggerName", "ZeroLogger").
			Logger()

		zerolog.SetGlobalLevel(l.getLogLevel())

		zeroSinLogger = &logger
	})
	l.logger = zeroSinLogger
}
func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Debug().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.
		Debug().
		Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Info().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.
		Info().
		Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.
		Info().
		Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Warn().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.
		Info().
		Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	l.logger.
		Fatal().
		Str("Category", string(cat)).
		Str("SubCategory", string(sub)).
		Fields(logParamsToZeroParams(extra)).
		Msg(msg)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.
		Info().
		Msgf(template, args...)
}
