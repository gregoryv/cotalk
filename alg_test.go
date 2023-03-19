package cotalk

import (
	"testing"
)

func BenchmarkAlg1(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg1); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg2(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg2); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg3(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg3); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg4(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg4); err != nil {
			b.Fatal(err)
		}
	}
}
