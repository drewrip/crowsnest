#!/bin/bash

./benchmark -i
for i in {1..100}
do
	for nn in {1..51..2}
	do
	./benchmark -n $nn
	done
done