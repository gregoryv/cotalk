package ex30

import (
	"testing"

	"github.com/preferit/cotalk"
)

func TestRun(t *testing.T) {
	if err := cotalk.CheckAlgorithm(Run); err != nil {
		t.Error(err)
	}
}
