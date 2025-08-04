package common

type Lemma interface {
	New(randomCount int64, seed int64)

	Generate() []any
}

type Type string

// TODO: Add more types as needed
const (
	Int   Type = "int"
	Float Type = "float"

	Custom    Type = "custom"    // User-defined types
	Undefined Type = "undefined" // For testing purposes only
)
