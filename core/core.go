package core

import (
	"fmt"
	"testing"

	"github.com/chriso345/lemma/common"
	"github.com/chriso345/lemma/corollary"
)

type TestFunc func(any) bool

func ForAll(t testing.TB, g common.Type, f TestFunc, c_ ...corollary.Corollary) {
	t.Helper()

	c := corollary.DefaultCorollary()
	if len(c_) > 1 {
		panic("Maximum of one corollary allowed")
	} else if len(c_) == 1 {
		c = &c_[0]
	}

	var lemma common.Lemma
	switch g {
	case common.Int:
		lemma = &intLemma{}
	case common.Float:
		lemma = &floatLemma{}
	case common.Custom:
		if c.Custom == nil {
			panic("Custom lemma must be provided for Custom type")
		}
		lemma = c.Custom
	default:
		panic(fmt.Sprintf("Unsupported type: %s", g))
	}

	lemma.New(c.RandomCount, c.Seed)
	generated := lemma.Generate()
	for _, item := range generated {
		if !f(item) {
			t.Errorf("ForAll failed for item: %v", item)
		}
	}
}
