package logger

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var reader, writer *os.File
var fakeLoggerInitialized bool = false

func initLoggers(t *testing.T) (*os.File, *os.File) {
	var err error
	if !fakeLoggerInitialized {
		reader, writer, err = os.Pipe()
		if err != nil {
			assert.Fail(t, "couldn't get os Pipe: %v", err)
		}
		InitLoggers(writer)
		fakeLoggerInitialized = true
	}
	return reader, writer
}

func getLastMessage(t *testing.T) (message string) {
	if !fakeLoggerInitialized {
		assert.Fail(t, "couldn't get os Pipe: %v", nil)
	}
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	return scanner.Text()
}

func getLastMessageJson(t *testing.T) (message map[string]interface{}) {
	msg := getLastMessage(t)
	json.Unmarshal([]byte(msg), &message)
	return
}

func TestWarn(t *testing.T) {
	initLoggers(t)
	Warn("foobar")
	msg := getLastMessageJson(t)
	assert.Contains(t, msg, "message")
	assert.Contains(t, msg, "level")
	assert.Contains(t, msg, "logger")
	assert.Equal(t, "foobar", msg["message"])
	assert.Equal(t, "warn", msg["level"])
	assert.Equal(t, "debug", msg["logger"])
}

func TestError(t *testing.T) {
	initLoggers(t)
	err := errors.New("test error, can discard")
	Error(err, "barbaz")
	msg := getLastMessageJson(t)
	assert.Contains(t, msg, "message")
	assert.Contains(t, msg, "logger")
	assert.Contains(t, msg, "level")
	assert.Contains(t, msg, "error")
	assert.Equal(t, "barbaz", msg["message"])
	assert.Equal(t, "error", msg["level"])
	assert.Equal(t, "test error, can discard", msg["error"])
	assert.Equal(t, "error", msg["logger"])
}

func TestAccessLogger(t *testing.T) {
	initLoggers(t)
	GetAccessLogger().Info().Msg("test")
	msg := getLastMessageJson(t)
	assert.Contains(t, msg, "time")
	assert.Contains(t, msg, "logger")
	assert.Contains(t, msg, "message")
	assert.Contains(t, msg, "level")
	assert.Equal(t, "test", msg["message"])
	assert.Equal(t, "access-log", msg["logger"])
}
