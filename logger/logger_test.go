package logger

import (
	"bufio"
	"encoding/json"
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
	assert.Equal(t, "foobar", msg["message"])
	assert.Equal(t, "warn", msg["level"])
}
