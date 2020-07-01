package utils_test

import (
	"gnt-cc/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInSlice(t *testing.T) {
	var stringSlice = []string{"this", "that", "these", "those"}

	assert.True(t, utils.IsInSlice("this", stringSlice))
	assert.False(t, utils.IsInSlice("not_in_slice", stringSlice))
}
