package cotalk

import (
	"fmt"
	"testing"
	"time"
)

func BenchmarkX(b *testing.B) {
	fmt.Println("N is", b.N)
	for i := 0; i < b.N; i++ {
		X()
	}
}

func X() {
	time.Sleep(time.Millisecond)
}
