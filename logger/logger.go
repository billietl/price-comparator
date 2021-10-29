package logger

import (
	"fmt"
	"io"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

var accessLogger zerolog.Logger
var errorLogger zerolog.Logger
var debugLogger zerolog.Logger

func InitLoggers(output io.Writer) {
	wr := diode.NewWriter(output, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})

	baseLogger := zerolog.New(output).With().Timestamp().Logger()
	delayedLogger := zerolog.New(wr).With().Timestamp().Logger()

	accessLogger = delayedLogger.With().
		Str("logger", "access-log").
		Logger()
	debugLogger = delayedLogger.With().
		Str("logger", "debug").
		Logger()
	errorLogger = baseLogger.With().
		Str("logger", "error").
		Logger()
}

func GetAccessLogger() *zerolog.Logger {
	return &accessLogger
}

func Error(err error, msg string) {
	errorLogger.Err(err).Msg(msg)
}

func Warn(msg string) {
	debugLogger.Warn().Msg(msg)
}

func Info(msg string) {
	debugLogger.Info().Msg(msg)
}
