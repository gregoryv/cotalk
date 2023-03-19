package cotalk

import (
	"testing"
	"time"
)

func BenchmarkX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		X()
	}
}

func X() {
	time.Sleep(time.Millisecond)
}
