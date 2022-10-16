#!/bin/sh
echo reset halt | nc -N localhost 4444
gdb-multiarch --batch -x gdb_target "$@" -x 'x/i $pc'

