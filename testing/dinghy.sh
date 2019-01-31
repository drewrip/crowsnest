#!/bin/bash

#./dinghy -i -d "dinghy"
for i in {1..5}
do
	for nn in {3..51..2}
	do
	./dinghy -n $nn -r -d "dinghy"
	done
done