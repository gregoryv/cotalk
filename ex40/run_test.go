package ex40

import (
	"testing"

	. "github.com/preferit/cotalk"
)

func TestRun(t *testing.T) {
	srv := SetupServer()
	defer srv.Close()
	if err := NewLettersProblem().Solve(srv, Run); err != nil {
		t.Error(err)
	}
}
