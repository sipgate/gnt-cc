package utils

import (
	"gnt-cc/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsvalidCluster(t *testing.T) {
	config.Parse()
	validClusterName := "production-cluster"
	assert.True(t, IsValidCluster(validClusterName))

	invalidClusterName := "myInvalidCluster"
	assert.False(t, IsValidCluster(invalidClusterName))
}
