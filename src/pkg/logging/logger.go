package logging

import "github.com/AmirHosein-Kahrani/Car-Center-Web/config"

type Logger interface {
	init()

	// Debug
	Debug(cat Category, subSubCategory, message string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})
	// Info
	Info(cat Category, subSubCategory, message string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})
	// Warn
	Warn(cat Category, subSubCategory, message string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})
	// Error
	Error(err error, cat Category, subSubCategory, message string, extra map[ExtraKey]interface{})
	Errorf(err error, template string, args ...interface{})
	// Panic
	Fatal(err error, cat Category, subSubCategory, message string, extra map[ExtraKey]interface{})
	Fatalf(err error, template string, args ...interface{})
}

func NewLogger(cfg *config.Config,) *Logger {

	return nil
}



//  file <- filebeat -> elasticSearch -> Kibana