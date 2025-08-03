package core

type Lemma interface {
	Generate(string) any // TODO: Change string type to a category type?
	All() []any
	Some() []any

	Number(int) any
	Type() Type
}

type intLemma struct {
}

var _ Lemma = (*intLemma)(nil)

func (l *intLemma) Generate(s string) any {
	return 0 // TODO: Implement actual logic
}

func (l *intLemma) All() []any {
	ints := []int{1, 2, 3} // TODO: Implement actual logic
	anys := make([]any, len(ints))
	for i, v := range ints {
		anys[i] = v
	}
	return anys
}

func (l *intLemma) Some() []any {
	ints := []int{1, 2} // TODO: Implement actual logic
	anys := make([]any, len(ints))
	for i, v := range ints {
		anys[i] = v
	}
	return anys
}

func (l *intLemma) Number(i int) any {
	return 3 // TODO: Implement actual logic
}

func (l *intLemma) Type() Type {
	return Int // TODO: Implement actual logic
}

type Type string

// TODO: Add more types as needed
const (
	Int   Type = "int"
	Float Type = "float"
)
