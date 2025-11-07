package core

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/chriso345/gore/assert"
)

func TestExpandArg_SingleDigit(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%d", r)
	assert.Equal(t, len(vals), 1)
	_, err := strconv.Atoi(vals[0])
	assert.Nil(t, err)
}

func TestExpandArg_MultipleDigits(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%d5", r)
	assert.Equal(t, len(vals), 5)
	for _, v := range vals {
		_, err := strconv.Atoi(v)
		assert.Nil(t, err)
	}
}

func TestExpandArg_DigitsWithRange(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%d3:5:7", r)
	assert.Equal(t, len(vals), 3)
	for _, v := range vals {
		n, _ := strconv.Atoi(v)
		assert.GreaterOrEqual(t, n, 5)
		assert.LessOrEqual(t, n, 7)
	}
}

func TestExpandArg_SingleFloat(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%f", r)
	assert.Equal(t, len(vals), 1)
	_, err := strconv.ParseFloat(vals[0], 64)
	assert.Nil(t, err)
}

func TestExpandArg_MultipleFloats(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%f3", r)
	assert.Equal(t, len(vals), 3)
	for _, v := range vals {
		_, err := strconv.ParseFloat(v, 64)
		assert.Nil(t, err)
	}
}

func TestExpandArg_FloatsWithRange(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%f2:1.5:3.5", r)
	assert.Equal(t, len(vals), 2)
	for _, v := range vals {
		f, _ := strconv.ParseFloat(v, 64)
		assert.GreaterOrEqual(t, f, 1.5)
		assert.LessOrEqual(t, f, 3.5)
	}
}

func TestExpandArg_LiteralString(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("literal", r)
	assert.Equal(t, len(vals), 1)
	assert.Equal(t, vals[0], "literal")
}

func TestExpandArg_IntEdgeCases(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%d4:0:0", r)
	assert.Equal(t, len(vals), 4)
	for _, v := range vals {
		n, _ := strconv.Atoi(v)
		assert.Equal(t, n, 0)
	}

	vals = expandArg("%d3:9:9", r)
	for _, v := range vals {
		n, _ := strconv.Atoi(v)
		assert.Equal(t, n, 9)
	}
}

func TestExpandArg_FloatEdgeCases(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	vals := expandArg("%f3:2.5:2.5", r)
	assert.Equal(t, len(vals), 3)
	for _, v := range vals {
		f, _ := strconv.ParseFloat(v, 64)
		assert.Equal(t, f, 2.5)
	}
}
