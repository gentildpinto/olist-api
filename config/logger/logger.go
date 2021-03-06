package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Initialize(version string) (err error) {
	config := zap.NewProductionConfig()

	if os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true" {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))

	if err != nil {
		return
	}

	log = &Logger{
		logger:  logger,
		version: version,
	}

	return
}

func Log(message interface{}) (err error) {
	if log == nil {
		err = errors.New("logger not configured")
		fmt.Println(err)
		panic(err)
	}

	switch messageType := message.(type) {
	case nil:
		return
	case error:
		log.logger.Error(messageType.Error())
		return messageType
	case string:
		log.logger.Info(messageType)
	default:
		if JSON(message) != "" {
			log.logger.Info(JSON(message))
		}
	}

	return
}

func JSON(value interface{}) string {
	bytes, err := json.MarshalIndent(value, "", "  ")

	if err != nil {
		return ""
	}

	return string(bytes)
}

type Logger struct {
	logger  *zap.Logger
	version string
}

var log *Logger
