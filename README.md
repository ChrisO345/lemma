# lemma

`lemma` is a simple Go package for generating dynamic test data for unit tests, inspired by property-based testing libraries like Hypothesis.

---

## Features

- **Randomized Data Generation** - Generate integers, floats, and custom command-based test data.
- **Edge Case Coverage** - Includes extreme values for numeric types to help uncover edge case bugs.
- **Custom Lemmas** - Users can define their own generators for complex types.
- **Command Integration** - Dynamically generate arguments and execute commands for integration testing.

---

## Installation

`lemma` is available on GitHub and can be installed using Go modules:

```bash
go get github.com/chriso345/lemma
````

---

## Usage

### Basic Usage

Import `lemma` and generate simple test data:

```go
import (
  "testing"

  "github.com/chriso345/lemma"
)

func TestIntegers(t *testing.T) {
  lemma.Test(t, lemma.Int, func(x any) bool {
    n := x.(int)
    return n*2%2 == 0
  })
}
```

### Command-Based Lemmas

Generate test commands dynamically, optionally using placeholders like `%d` for integers or `%f` for floats:

```go
cmd := lemma.Command("echo", "%d3", "%f2", "literal")
lemma.Test(t, cmd, func(result any) bool {
  value := result.(lemma.CommandResult)
  args := value.Args
  res := value.Result
  // args is ["echo", "5", "3", "7", "1.234", "0.987", "literal"] (example)
  return len(args) > 0
})
```

### Custom Lemmas

Define your own generator for custom types:

```go
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

func TestCustomLemma(t *testing.T) {
	customCorollary := corollary.DefaultCorollary()
	customCorollary.Custom = &CustomLemma{}

  lemma.Test(mt, lemma.Custom, func(x any) bool {
    value := x.(string)
    return value == ""
  }, *customCorollary)
}
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
