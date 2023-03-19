package ex10

import (
	"testing"

	. "github.com/preferit/cotalk"
)

func TestRun(t *testing.T) {
	srv := SetupServer(t)
	if err := NewLettersProblem().Solve(srv, Run); err != nil {
		t.Error(err)
	}
}
