#!/bin/bash

for n in {1..51..2}
do
	../../kvdinghy/kvdinghy --test -s $n &>/dev/null &
	disown
	sleep 2
	start=($SECONDS)
	for i in {1..1000}
	do
		curl -d '{"key": "k'$i'", "value": '$i'}' -X POST http://localhost:8000/set
		curl -d '{"key": "k'$i'"}' -X GET http://localhost:8000/get
		echo ""
	done
	echo $n" "$((SECONDS-start)) >> kv.dat
	killall kvdinghy
done