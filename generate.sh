#!/bin/bash -e
pushd ex07; go test -benchmem -bench . | gocolor > bench_result.txt; popd
