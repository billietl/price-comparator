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
	accessLogger = zerolog.New(wr).With().
		Timestamp().
		Logger()
	debugLogger = zerolog.New(wr).With().Logger()
	errorLogger = zerolog.New(output).With().Logger()
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
