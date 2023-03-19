#!/bin/bash -e


go test -count 1 -v -run TestSequential . | gocolor | aha -n > testdata/sequential_test.html

case $1 in
    ex10|ex20|ex30|ex40|ex50)
	pushd $1; go test -count 1 -v . | gocolor | aha -n > test_result.html; popd
	;;
    ex70)
	pushd $1; go test -benchmem -bench . | gocolor | aha -n > bench_result.html; popd
	;;
esac
