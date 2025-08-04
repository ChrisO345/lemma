package core

import (
	"math"
	"math/rand"
	"time"

	"github.com/chriso345/lemma/common"
)

type floatLemma struct {
	randomCount int64
	seed        int64
}

func (l *floatLemma) New(randomCount int64, seed int64) {
	l.randomCount = randomCount
	l.seed = seed
}

func (l *floatLemma) Generate() []any {
	edgeCases := []float64{
		math.Inf(-1), -math.MaxFloat64, -1000.0, -100.0, -10.0, -1.0, 0.0, 1.0, 10.0, 100.0, 1000.0, math.MaxFloat64, math.Inf(1),
	}

	if l.seed == -1 {
		l.seed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(l.seed))

	for i := int64(0); i < l.randomCount; i++ {
		edgeCases = append(edgeCases, r.Float64()*math.MaxFloat64)
	}

	anys := make([]any, len(edgeCases))
	for i, v := range edgeCases {
		anys[i] = v
	}
	return anys
}

// Force implementation of Lemma interface
var _ common.Lemma = (*floatLemma)(nil)
