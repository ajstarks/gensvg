#!/bin/sh
list=$(cat rl)
cd cmd
for i in $list
do
	cd $i
	google-chrome $i.svg 2>/dev/null &
	echo -n "$i "
	cd ..
done
echo
