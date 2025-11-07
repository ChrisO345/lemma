package lemma_test

import (
	"fmt"
	"strings"
	"testing"

	// "github.com/chriso345/gore/assert"
	"github.com/chriso345/gore/assert"
	"github.com/chriso345/lemma"
	"github.com/chriso345/lemma/corollary"
)

func TestLemma_TestInt(t *testing.T) {
	mt := newMockT()

	lemma.Test(mt, lemma.Int, func(x any) bool {
		d := x.(int) * 2
		return d%2 == 0
	})

	if mt.failed {
		t.Error("Expected test to pass, but it failed")
	}

	mt = newMockT()

	lemma.Test(mt, lemma.Int, func(x any) bool {
		d := x.(int) * 2
		return d%2 == 1
	})

	if !mt.failed {
		t.Error("Expected test to fail, but it passed")
	}
}

func TestLemma_TestFloat(t *testing.T) {
	mt := newMockT()

	lemma.Test(mt, lemma.Float, func(x any) bool {
		if x.(float64) <= 1 {
			return true
		}

		d := x.(float64) * 2
		return d > 0
	})

	if mt.failed {
		t.Error("Expected test to pass, but it failed")
	}

	mt = newMockT()

	lemma.Test(mt, lemma.Float, func(x any) bool {
		d := x.(float64) * -1
		return d < 0
	})

	if !mt.failed {
		fmt.Println(mt.logs)
		t.Error("Expected test to fail, but it passed")
	}
}

func TestLemma_TestUndefinedType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for unsupported type, but did not panic")
		}
	}()

	mt := newMockT()

	lemma.Test(mt, lemma.Undefined, func(x any) bool {
		return true
	})

	if !mt.failed {
		t.Error("Expected test to fail due to unsupported type, but it passed")
	}
}

func TestLemma_TestWithCustomLemma(t *testing.T) {
	mt := newMockT()

	customCorollary := corollary.DefaultCorollary()
	customCorollary.Custom = &CustomLemma{}
	customCorollary.RandomCount = 5

	lemma.Test(mt, lemma.Custom, func(x any) bool {
		value := x.(string)
		return value != ""
	}, *customCorollary)

	if mt.failed {
		t.Error("Expected test to pass with custom lemma, but it failed")
	}

	lemma.Test(mt, lemma.Custom, func(x any) bool {
		value := x.(string)
		return value == ""
	}, *customCorollary)

	assert.Equal(t, len(mt.logs), 5)
	if !mt.failed {
		t.Error("Expected test to fail with custom lemma, but it passed")
	}
}

func TestLemma_TestWithCustomLemmaNotProvided(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for missing custom lemma, but did not panic")
		}
	}()

	mt := newMockT()

	lemma.Test(mt, lemma.Custom, func(x any) bool {
		return true
	})
}

// === COMMAND LEMMA ===
func TestLemma_TestCommandLemma(t *testing.T) {
	mt := newMockT()

	commandCorollary := corollary.DefaultCorollary()
	commandCorollary.RandomCount = 5

	lemma.Test(mt, lemma.Command("echo", "%d10:5:15"), func(x any) bool {
		value := x.(lemma.CommandResult)
		a := value.Args[1:] // Exclude command itself
		b := strings.Split(strings.TrimSpace(value.Result), " ")

		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

		return true
	}, *commandCorollary)

	if mt.failed {
		t.Error("Expected test to pass with command lemma, but it failed")
	}
}

// === CUSTOM LEMMA ===
type CustomLemma struct {
	randomCount int64
	seed        int64
}

func (c *CustomLemma) New(randomCount int64, seed int64) {
	c.randomCount = randomCount
	c.seed = seed
}

func (c *CustomLemma) Generate() []any {
	anys := make([]any, c.randomCount)

	for i := int64(0); i < c.randomCount; i++ {
		anys[i] = fmt.Sprintf("custom_value_%d", i)
	}
	return anys
}

var _ = (*CustomLemma)(nil)

// === MOCK TEST ===
type mockT struct {
	testing.TB
	failed bool
	logs   []string
}

func (m *mockT) Helper() {}
func (m *mockT) Errorf(format string, args ...any) {
	m.failed = true
	m.logs = append(m.logs, format)
}

func newMockT() *mockT {
	return &mockT{}
}
