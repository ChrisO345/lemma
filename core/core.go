package core

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	"github.com/chriso345/lemma/common"
	"github.com/chriso345/lemma/corollary"
)

type TestFunc func(any) bool

func ForAll(t testing.TB, g LemmaType, f TestFunc, c_ ...corollary.Corollary) {
	t.Helper()

	c := corollary.DefaultCorollary()
	if len(c_) > 1 {
		panic("Maximum of one corollary allowed")
	} else if len(c_) == 1 {
		c = &c_[0]
	}

	var lemma common.Lemma
	if g.TypeName == "custom" {
		if c.Custom == nil {
			panic("Custom lemma must be provided for Custom type")
		}
		lemma = c.Custom
	} else {
		lemma = g.Factory()

		if g.Factory == nil {
			panic("Lemma factory function is nil")
		}
	}

	lemma.New(c.RandomCount, c.Seed)
	generated := lemma.Generate()
	if g.TypeName == "command" {
		runCommandGenerations(t, generated, f)
		return
	}
	for _, item := range generated {
		if !f(item) {
			t.Errorf("ForAll failed for item: %v", item)
		}
	}
}

func runCommandGenerations(t testing.TB, generated []any, f TestFunc) {
	t.Helper()

	for _, item := range generated {
		cmdArgs, ok := item.([]string)
		if !ok {
			t.Errorf("Expected []string for command lemma, got: %T", item)
			continue
		}

		if len(cmdArgs) == 0 {
			t.Errorf("Empty command in command lemma")
			continue
		}

		// Use the first element as command, the rest as args
		cmdName := cmdArgs[0]
		args := cmdArgs[1:]

		// Prepare command using sh -c for safety
		cmd := exec.Command("sh", "-c", cmdName+" "+joinArgs(args))

		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out

		err := cmd.Run()
		resultStr := out.String()
		if err != nil {
			resultStr = fmt.Sprintf("ERROR: %v\n%s", err, resultStr)
		}

		// Wrap in CommandResult for the property function
		ar := common.CommandResult{
			Args:   cmdArgs,
			Result: resultStr,
		}

		if !f(ar) {
			t.Errorf("ForAll failed for command args: %v\nOutput:\n%s", cmdArgs, resultStr)
		}
	}
}

// joinArgs safely joins args for sh -c execution
func joinArgs(args []string) string {
	escaped := make([]string, len(args))
	for i, a := range args {
		// simple quoting to handle spaces and special chars
		escaped[i] = fmt.Sprintf("%q", a)
	}
	return fmt.Sprintf("%s", stringJoin(escaped, " "))
}

// stringJoin is a lightweight strings.Join alternative
func stringJoin(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for _, s := range strs[1:] {
		res += sep + s
	}
	return res
}
