package lemma

import (
	"testing"

	"github.com/chriso345/lemma/internal/core"
)

func TestLemmaPass(t *testing.T) {
	core.ForAll(t, core.Int, func(x any) bool {
		d := x.(int) * 2
		return d%2 == 0
	})
}

func TestLemmaFail(t *testing.T) {
	core.ForAll(t, core.Int, func(x any) bool {
		d := x.(int) * 2
		return d%2 == 1 // This will fail for even numbers
	})
}
