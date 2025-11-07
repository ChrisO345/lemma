package core

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/chriso345/lemma/common"
)

func Command(command string, args ...string) LemmaType {
	return LemmaType{
		TypeName: "command",
		Factory: func() common.Lemma {
			return &commandLemma{
				command: command,
				args:    args,
			}
		},
	}
}

type commandLemma struct {
	randomCount int64
	seed        int64
	command     string
	args        []string
}

func (l *commandLemma) New(randomCount int64, seed int64) {
	l.randomCount = randomCount
	l.seed = seed
}

func (l *commandLemma) Generate() []any {
	r := rand.New(rand.NewSource(l.seed))
	results := make([]any, 0, l.randomCount)

	for i := int64(0); i < l.randomCount; i++ {
		cmdArgs := []string{l.command}
		for _, arg := range l.args {
			expanded := expandArg(arg, r)
			cmdArgs = append(cmdArgs, expanded...)
		}
		results = append(results, cmdArgs)
	}

	return results
}

// expandArg handles %d, %f, and optional counts and ranges
func expandArg(arg string, r *rand.Rand) []string {
	// Determine type (%d or %f)
	if strings.HasPrefix(arg, "%d") || strings.HasPrefix(arg, "%f") {
		isFloat := strings.HasPrefix(arg, "%f")
		rest := arg[2:] // strip %d or %f
		count := 1
		min := 0.0
		max := 0.0
		if isFloat {
			max = math.MaxFloat64
		} else {
			max = 9
		}

		// parse count:min:max
		if len(rest) > 0 {
			parts := strings.Split(rest, ":")
			// count
			if n, err := strconv.Atoi(parts[0]); err == nil {
				count = n
			}
			// min
			if len(parts) > 1 {
				if val, err := strconv.ParseFloat(parts[1], 64); err == nil {
					min = val
				}
			}
			// max
			if len(parts) > 2 {
				if val, err := strconv.ParseFloat(parts[2], 64); err == nil {
					max = val
				}
			}
		}

		// generate random values
		values := make([]string, count)
		for i := 0; i < count; i++ {
			if isFloat {
				v := min + r.Float64()*(max-min)
				values[i] = fmt.Sprintf("%f", v)
			} else {
				v := int(min) + r.Intn(int(max-min+1))
				values[i] = strconv.Itoa(v)
			}
		}
		return values
	}

	// literal string
	return []string{arg}
}

// Force implementation of Lemma interface
var _ common.Lemma = (*commandLemma)(nil)
