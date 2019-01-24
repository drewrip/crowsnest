#!/bin/bash

./dinghy -i
for i in {1..200}
do
	for nn in {3..51..2}
	do
	./dinghy -n $nn
	done
done