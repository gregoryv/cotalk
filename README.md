Presentation about concurrency design in Go

## Quick start

To build the presentation

  $ cd docs
  $ go run .
  $ $BROWSER index.html

the testdata fragments where generated with eg.

  $ go test -benchmem -bench=BenchmarkAlg01 . | gocolor | aha -n > testdata/alg01_bench.html

