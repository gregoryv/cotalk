#!/bin/bash -e
pushd ex07; go test -benchmem -bench . | gocolor | aha -n > bench_result.html; popd
