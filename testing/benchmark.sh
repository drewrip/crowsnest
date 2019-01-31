#!/bin/bash

#./benchmark -i -d "benchmark"
for i in {1..5}
do
	for nn in {3..51..2}
	do
	./benchmark -n $nn -d "benchmark"
	done
done