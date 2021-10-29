package logger

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var reader, writer *os.File
var scanner *bufio.Scanner

func init() {
	var err error
	reader, writer, err = os.Pipe()
	if err != nil {
		log.Fatalf("couldn't get os Pipe: %v", err)
	}
	InitLoggers(writer)
	scanner = bufio.NewScanner(reader)
}

func getLastMessageJson() (message map[string]interface{}) {
	scanner.Scan()
	json.Unmarshal(scanner.Bytes(), &message)
	return
}

func TestWarn(t *testing.T) {
	Warn("foobar")
	msg := getLastMessageJson()
	assert.Contains(t, msg, "time")
	assert.Contains(t, msg, "message")
	assert.Contains(t, msg, "level")
	assert.Contains(t, msg, "logger")
	assert.Equal(t, "foobar", msg["message"])
	assert.Equal(t, "warn", msg["level"])
	assert.Equal(t, "debug", msg["logger"])
}

func TestInfo(t *testing.T) {
	Info("hugre")
	msg := getLastMessageJson()
	assert.Contains(t, msg, "time")
	assert.Contains(t, msg, "message")
	assert.Contains(t, msg, "level")
	assert.Contains(t, msg, "logger")
	assert.Equal(t, "hugre", msg["message"])
	assert.Equal(t, "info", msg["level"])
	assert.Equal(t, "debug", msg["logger"])
}

func TestError(t *testing.T) {
	err := errors.New("test error, can discard")
	Error(err, "barbaz")
	msg := getLastMessageJson()
	assert.Contains(t, msg, "time")
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
	GetAccessLogger().Info().Msg("test")
	msg := getLastMessageJson()
	assert.Contains(t, msg, "time")
	assert.Contains(t, msg, "logger")
	assert.Contains(t, msg, "message")
	assert.Contains(t, msg, "level")
	assert.Equal(t, "test", msg["message"])
	assert.Equal(t, "access-log", msg["logger"])
}
