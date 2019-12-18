#!/bin/sh

if [ "x" = "x$1" ];
then
	echo "usage: $0 <file.bin>"
	exit 1
fi

echo reset halt | nc -N localhost 4444
sleep .1
echo flash write_image erase "$1" 0x8000000 | nc -N localhost 4444
sleep .1
( echo verify_image "$1" 0x8000000 ; echo reset run ) | nc -N localhost 4444

echo
