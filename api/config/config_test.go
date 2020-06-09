package config

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestParseLogLevel(t *testing.T) {
	log.StandardLogger().ExitFunc = nil

	if parseLogLevel("debug") != log.DebugLevel {
		t.Errorf("failed to parse 'debug' loglevel")
	}

	if parseLogLevel("info") != log.InfoLevel {
		t.Errorf("failed to parse 'info' loglevel")
	}

	if parseLogLevel("warning") != log.WarnLevel {
		t.Errorf("failed to parse 'warning' loglevel")
	}

	if parseLogLevel("error") != log.ErrorLevel {
		t.Errorf("failed to parse 'error' loglevel")
	}

	if parseLogLevel("fatal") != log.FatalLevel {
		t.Errorf("failed to parse 'fatal' loglevel")
	}

	if parseLogLevel("randomString") != log.WarnLevel {
		t.Errorf("failed to set fallback 'warning' loglevel")
	}

}

func TestIsInSlice(t *testing.T) {
	var stringSlice = []string{"this", "that", "these", "those"}

	if !isInSlice("this", stringSlice) {
		t.Errorf("sample string 'this' falsly not detected in slice")
	}

	if isInSlice("not_in_slice", stringSlice) {
		t.Errorf("sample string 'not_in_slice' falsly detected slice")
	}
}

func TestGetClusterConfig(t *testing.T) {
	Parse()
	// the following function either returns a ClusterConfig struct or panics (and hence should be improved)
	GetClusterConfig("production-cluster")
}
