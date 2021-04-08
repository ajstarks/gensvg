#!/bin/sh
list=$(cat rl)
cd cmd
for i in $list
do
	cd $i
	go run $i.go > $i.svg
	echo -n "$i "
	cd ..
done
echo
