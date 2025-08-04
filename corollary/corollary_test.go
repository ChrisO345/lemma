package corollary

import (
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestDefaultCorollary(t *testing.T) {
	c := DefaultCorollary()
	assert.Equal(t, c.Seed, -1)
	assert.Equal(t, c.RandomCount, 100)
	assert.Nil(t, c.Custom)
}
