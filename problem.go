package cotalk

import (
	"net/http"
)

// Problem defines the thing that needs to be solved
type Problem interface {
	// Solve returns an error if the given algorithm does not solve
	// the problem.
	Solve(Algorithm) error
}

// Algorithm is any func that takes some work and returns it's result.
// The validity of the result must be verified outside.
type Algorithm func(work []*http.Request) (result []*http.Response)
