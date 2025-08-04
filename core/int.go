package core

import (
	"math"
	"math/rand"
	"time"

	"github.com/chriso345/lemma/common"
)

type intLemma struct {
	randomCount int64
	seed        int64
}

func (l *intLemma) New(randomCount int64, seed int64) {
	l.randomCount = randomCount
	l.seed = seed
}

func (l *intLemma) Generate() []any {
	edgeCases := []int{
		math.MinInt, -math.MaxInt, -1000, -100, -10, -1, 0, 1, 10, 100, 1000, math.MaxInt,
	}

	if l.seed == -1 {
		l.seed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(l.seed))

	for i := int64(0); i < l.randomCount; i++ {
		edgeCases = append(edgeCases, r.Intn(math.MaxInt))
	}

	anys := make([]any, len(edgeCases))
	for i, v := range edgeCases {
		anys[i] = v
	}
	return anys
}

// Force implementation of Lemma interface
var _ common.Lemma = (*intLemma)(nil)
