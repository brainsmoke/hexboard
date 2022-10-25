#!/bin/bash
hexfile="$1"
openocd -f openocd_k20.cfg -c 'init;halt;program '"$hexfile"' verify'
