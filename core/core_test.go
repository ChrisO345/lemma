package core

import (
	"math"
	"testing"

	"github.com/chriso345/gore/assert"
)

var LENGTH int64 = 10
var SEED int64 = 5

func Test_Int(t *testing.T) {
	edgeCases := []any{
		math.MinInt, -math.MaxInt, -1000, -100, -10, -1, 0, 1, 10, 100, 1000, math.MaxInt,
	}

	var l intLemma
	var m intLemma

	l.New(LENGTH, SEED)
	m.New(LENGTH, SEED)

	lGen := l.Generate()
	mGen := m.Generate()

	assert.Equal(t, l, m)
	assert.Length(t, lGen, int(LENGTH)+len(edgeCases))
	assert.Length(t, mGen, int(LENGTH)+len(edgeCases))

	if len(lGen) == len(mGen) {
		for idx := range len(lGen) {
			if idx < len(edgeCases) {
				assert.Equal(t, lGen[idx], edgeCases[idx])
			}
			assert.Equal(t, lGen[idx], mGen[idx])
		}
	}
}

func Test_Float(t *testing.T) {
	edgeCases := []any{
		math.Inf(-1), -math.MaxFloat64, -1000.0, -100.0, -10.0, -1.0, 0.0, 1.0, 10.0, 100.0, 1000.0, math.MaxFloat64, math.Inf(1),
	}

	var l floatLemma
	var m floatLemma

	l.New(LENGTH, SEED)
	m.New(LENGTH, SEED)

	lGen := l.Generate()
	mGen := m.Generate()

	assert.Equal(t, l, m)
	assert.Length(t, lGen, int(LENGTH)+len(edgeCases))
	assert.Length(t, mGen, int(LENGTH)+len(edgeCases))

	if len(lGen) == len(mGen) {
		for idx := range len(lGen) {
			if idx < len(edgeCases) {
				assert.Equal(t, lGen[idx], edgeCases[idx])
			}
			assert.Equal(t, lGen[idx], mGen[idx])
		}
	}
}
