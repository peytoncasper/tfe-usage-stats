package log

import (
	"go.uber.org/zap"
	"log"
	"os"
)

type logger struct {
	zap   *zap.Logger
	level zap.AtomicLevel
}

var wrappedLogger *logger

func init() {

	var config zap.Config

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		config = zap.NewProductionConfig()
	} else if logLevel == "DEBUG" {
		config = zap.NewDevelopmentConfig()
	}


	l, err := config.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		log.Fatal("error setting up logging: ", err)
	}

	wrappedLogger = &logger{
		zap:   l.Named("tfe-usage-stats"),
		level: config.Level,
	}
	defer wrappedLogger.zap.Sync()
}

func Fatal(msg string, fields ...zap.Field) {
	wrappedLogger.zap.Fatal(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	wrappedLogger.zap.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	wrappedLogger.zap.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	wrappedLogger.zap.Error(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	wrappedLogger.zap.Warn(msg, fields...)
}
