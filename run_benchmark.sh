#!/bin/bash

COUNT=5
BENCHMEM=true
GCCFLAGS='-O3 -march=native'
CURDIR=${PWD}

# Turn garbage collection settings
# Default is 100, off is 0
cd benchmarks

GOGC=100
echo "Running Basic Benchmarks"
GOGC=$GOGC go test go_basic_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="-O1 -march=native" -gcflags="-l" > $CURDIR/results/gc/basic_go
sleep 60
GOGC=$GOGC go test go_basic_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="-O1 -march=native" -gcflags="-l"> $CURDIR/results/gccgo/basic_go
sleep 60

GOGC=$GOGC go test pure_go_matrix_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/matrix
sleep 60
GOGC=$GOGC go test pure_go_matrix_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/matrix
sleep 60

GOGC=$GOGC go test math_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/math
sleep 60
GOGC=$GOGC go test math_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/math
sleep 60

GOGC=$GOGC go test rand_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/rand
sleep 60
GOGC=$GOGC go test rand_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/rand
sleep 60

GOGC=$GOGC go test strings_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/strings
sleep 60
GOGC=$GOGC go test strings_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/strings
sleep 60

GOGC=$GOGC go test encode_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/encoding
sleep 60
GOGC=$GOGC go test encode_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/encoding
sleep 60

GOGC=$GOGC go test bzip2_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/bzip2
sleep 60
GOGC=$GOGC go test bzip2_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/bzip2
sleep 60

GOGC=$GOGC go test md5_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/md5
sleep 60
GOGC=$GOGC go test md5_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/md5
sleep 60

GOGC=$GOGC go test search_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/search
sleep 60
GOGC=$GOGC go test search_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/search
sleep 60

GOGC=$GOGC go test writer_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/writer
sleep 60
GOGC=$GOGC go test writer_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/writer
sleep 60

GOGC=off
echo "Running Basic Benchmarks"
GOGC=$GOGC go test go_basic_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/basic_go
sleep 60
GOGC=$GOGC go test go_basic_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/basic_go
sleep 60

GOGC=$GOGC go test pure_go_matrix_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/matrix
sleep 60
GOGC=$GOGC go test pure_go_matrix_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/matrix
sleep 60

GOGC=$GOGC go test math_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/math
sleep 60
GOGC=$GOGC go test math_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/math
sleep 60

GOGC=$GOGC go test rand_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/rand
sleep 60
GOGC=$GOGC go test rand_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/rand
sleep 60

GOGC=$GOGC go test strings_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/strings
sleep 60
GOGC=$GOGC go test strings_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/strings
sleep 60

GOGC=$GOGC go test encode_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/encoding
sleep 60
GOGC=$GOGC go test encode_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/encoding
sleep 60

GOGC=$GOGC go test bzip2_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/bzip2
sleep 60
GOGC=$GOGC go test bzip2_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/bzip2
sleep 60

GOGC=$GOGC go test md5_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/md5
sleep 60
GOGC=$GOGC go test md5_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/md5
sleep 60

GOGC=$GOGC go test search_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/gcoff/search
sleep 60
GOGC=$GOGC go test search_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/gcoff/search
sleep 60

GOGC=$GOGC go test writer_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gc -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gc/writer
sleep 60
GOGC=$GOGC go test writer_test.go -count=$COUNT -bench=. -benchmem=$BENCHMEM -compiler=gccgo -gccgoflags="${GCCFLAGS}" > $CURDIR/results/gccgo/writer
sleep 60 
