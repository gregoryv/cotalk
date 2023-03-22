package cotalk

import (
	"context"
	"net/http"
	"testing"
	"time"
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

func BenchmarkAlg5(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg5); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg6(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg6); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg7(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg7); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg8(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()
	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg8); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg9(b *testing.B) {
	srv, problem := Setup()
	defer srv.Close()

	// wrap algorithm
	alg := func() Algorithm {
		return func(work []*http.Request) []*http.Response {
			ctx, _ := context.WithTimeout(context.Background(), 12*time.Millisecond)
			return Alg9(ctx, work)
		}
	}()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(alg); err != nil {
			b.Fatal(err)
		}
	}
}
