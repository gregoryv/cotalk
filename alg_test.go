package cotalk

import (
	"context"
	"net/http"
	"testing"
	"time"
)

const Letters = "0 1 2 3 4 5 6 7 8 9 a b c d e f"

func BenchmarkAlg01(b *testing.B) {
	// setup problem outside the loop
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg01); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg02(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg02); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg03(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg03); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg04(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg04); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg05(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg05); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg06(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg06); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg07(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg07); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg08(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(Alg08); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg09(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	// wrap algorithm
	alg := func() Algorithm {
		return func(work []*http.Request) []*http.Response {
			ctx, _ := context.WithTimeout(context.Background(), 12*time.Millisecond)
			return Alg09(ctx, work)
		}
	}()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(alg); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAlg10(b *testing.B) {
	problem := NewLetterChallenge(Letters)
	srv := problem.Server()
	defer srv.Close()
	b.ResetTimer()

	// wrap algorithm
	alg := func() Algorithm {
		return func(work []*http.Request) []*http.Response {
			ctx, _ := context.WithTimeout(context.Background(), 12*time.Millisecond)
			return Alg10(ctx, work)
		}
	}()

	for i := 0; i < b.N; i++ {
		if err := problem.Solve(alg); err != nil {
			b.Fatal(err)
		}
	}
}
