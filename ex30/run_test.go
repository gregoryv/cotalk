package ex30

import (
	"testing"

	"github.com/preferit/cotalk"
)

func TestRun(t *testing.T) {
	if err := cotalk.NewLettersProblem().Solve(Run); err != nil {
		t.Error(err)
	}
}
