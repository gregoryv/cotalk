#!/bin/bash -e


go test -benchmem -bench=BenchmarkAlg1 . | gocolor | aha -n > testdata/alg1_bench.html
go test -benchmem -bench=BenchmarkAlg2 . | gocolor | aha -n > testdata/alg2_bench.html
#go test -benchmem -bench=BenchmarkAlg3 . | gocolor | aha -n > testdata/alg3_bench.html
go test -benchmem -bench=BenchmarkAlg4 . | gocolor | aha -n > testdata/alg4_bench.html
go test -benchmem -bench=BenchmarkAlg5 . | gocolor | aha -n > testdata/alg5_bench.html
go test -benchmem -bench=BenchmarkAlg6 . | gocolor | aha -n > testdata/alg6_bench.html
go test -benchmem -bench=BenchmarkAlg7 . | gocolor | aha -n > testdata/alg7_bench.html
go test -benchmem -bench=BenchmarkAlg8 . | gocolor | aha -n > testdata/alg8_bench.html
