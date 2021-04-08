#!/bin/sh
list=$(cat cl)
cd cmd
for i in $list
do
	cd $i
	go build 
	echo -n "$i "
	cd ..
done
echo
