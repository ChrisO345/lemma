package core

import (
	"testing"
)

type TestFunc func(any) bool

func ForAll(t testing.TB, g Type, f TestFunc) {
	t.Helper()
	// FIXME: This is always using int types and should be decided based on the type of `g`.
	var lemma intLemma
	generated := lemma.All()
	for _, item := range generated {
		if !f(item) {
			t.Errorf("ForAll failed for item: %v", item)
		}
	}
}
