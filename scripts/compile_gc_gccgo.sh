#!/bin/bash

if [ "$#" -ne 2 ]; then
  echo "Usage: compile_both.sh [path/to/go_file.go] [path/to/output_file]"
  exit
fi

go build -o $2_gc $1
gccgo-8 $1 -o $2_gccgo
