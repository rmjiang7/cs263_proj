#!/bin/bash
for filename in gc/*; do
  benchstat gc/$(basename $filename) gccgo/$(basename $filename) >> benchstat_results
done
