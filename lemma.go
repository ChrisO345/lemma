package lemma

import (
	"github.com/chriso345/lemma/common"
	"github.com/chriso345/lemma/core"
)

// Test
var Test = core.ForAll

// === TYPES ===
var Int = core.Int
var Float = core.Float
var Command = core.Command
var Custom = core.Custom       // Custom type for user-defined lemmas
var Undefined = core.Undefined // Undefined type for error handling

// === Helpers ===
type CommandResult = common.CommandResult
