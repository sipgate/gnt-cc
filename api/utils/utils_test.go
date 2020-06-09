package utils

import (
	"gnt-cc/config"
	"testing"
)

func TestIsvalidCluster(t *testing.T) {
	config.Parse()
	validClusterName := "production-cluster"
	if !IsValidCluster(validClusterName) {
		t.Errorf("IsValidCluster failed to validate a proper cluster name")
	}

	invalidClusterName := "myInvalidCluster"
	if IsValidCluster(invalidClusterName) {
		t.Errorf("IsValidCluster failed to invalidate a wrong cluster name")
	}
}
