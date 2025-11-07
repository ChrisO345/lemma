package common

type Lemma interface {
	New(randomCount int64, seed int64)

	Generate() []any
}

type CommandResult struct {
	Args   []string
	Result string
}
