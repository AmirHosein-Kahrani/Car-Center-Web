package logging

import "github.com/AmirHosein-Kahrani/Car-Center-Web/config"

type Logger interface {
	Init()

	// Debug
	Debug(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})
	// Info
	Info(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})
	// Warn
	Warn(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})
	// Error
	Error(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})
	// Panic
	Fatal(cat Category, sub SubCategory, message string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {
	switch cfg.Logger.Logger {
	case "zap":
		return newZapLogger(cfg)
	case "zerolog":
		return newZeroLogger(cfg)
	}
	panic("logger not supported")
}

//  file <- filebeat -> elasticSearch -> Kibana
