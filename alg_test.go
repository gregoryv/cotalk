package cotalk

import (
	"testing"
)

func TestSequential(t *testing.T) {
	SolveLettersProblem(t, Sequential)
}

func TestConcurrentWaitGroup(t *testing.T) {
	SolveLettersProblem(t, ConcurrentWaitGroup)
}
