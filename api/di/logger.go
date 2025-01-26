package di

import (
	"teacher-grading-api/logger"
	"teacher-grading-api/shared/config"
)

func NewLogger(config *config.Configuration) (logger.Logger, error) {
	return logger.New(&logger.Option{
		Level:       logger.Level(config.LogLevel),
		Formatter:   logger.Formatter(config.LogFormatter),
		LogFilePath: config.LogFilePath,
		MaxSize:     config.LogMaxSize,
		MaxBackups:  config.LogMaxBackup,
		MaxAge:      config.LogMaxAge,
		Compress:    config.LogCompress,
	})
}
