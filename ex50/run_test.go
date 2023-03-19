package ex50

import (
	"testing"

	"github.com/preferit/cotalk"
)

func TestRun(t *testing.T) {
	if err := cotalk.Solve(Run); err != nil {
		t.Error(err)
	}
}
