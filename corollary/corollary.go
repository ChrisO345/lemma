package corollary

import "github.com/chriso345/lemma/common"

type Corollary struct {
	Seed        int64
	RandomCount int64
	Custom      common.Lemma
}

func DefaultCorollary() *Corollary {
	return &Corollary{
		Seed:        -1,  // Random Seed
		RandomCount: 100, // Number of random values to generate
		Custom:      nil, // Custom lemma pointer
	}
}
