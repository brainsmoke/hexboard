#!/bin/sh
echo reset halt | nc -N localhost 4444
gdb-multiarch -x gdb_target "$@"
