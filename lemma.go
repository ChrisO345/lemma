package lemma

import (
	"github.com/chriso345/lemma/common"
	"github.com/chriso345/lemma/core"
)

// Test
var Test = core.ForAll

// === TYPES ===
var Int = common.Int
var Float = common.Float
var Custom = common.Custom       // Custom type for user-defined lemmas
var Undefined = common.Undefined // Undefined type for error handling
