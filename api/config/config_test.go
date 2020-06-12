package config

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestParseLogLevel(t *testing.T) {
	log.StandardLogger().ExitFunc = nil

	assert.Equal(t, parseLogLevel("debug"), log.DebugLevel, "failed to parse loglevel 'debug'")
	assert.Equal(t, parseLogLevel("info"), log.InfoLevel, "failed to parse loglevel 'info'")
	assert.Equal(t, parseLogLevel("warning"), log.WarnLevel, "failed to parse loglevel 'warning'")
	assert.Equal(t, parseLogLevel("error"), log.ErrorLevel, "failed to parse loglevel 'error'")
	assert.Equal(t, parseLogLevel("fatal"), log.FatalLevel, "failed to parse loglevel 'fatal'")
	assert.Equal(t, parseLogLevel("randomString"), log.WarnLevel, "failed to parse illegal loglevel 'randomString'")
}

func TestIsInSlice(t *testing.T) {
	var stringSlice = []string{"this", "that", "these", "those"}

	assert.True(t, isInSlice("this", stringSlice))
	assert.False(t, isInSlice("not_in_slice", stringSlice))
}

func TestGetClusterConfig(t *testing.T) {
	Parse()
	// the following function either returns a ClusterConfig struct or panics (and hence should be improved)
	GetClusterConfig("production-cluster")
}
