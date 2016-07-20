#!/bin/bash
input="$(tail -1 /mnt/LONTAS/ExpControl/dire15//info/details.txt)"
set -f
array=(${input// / })
echo "${array[0]}"
#capshow /mnt/LONTAS/traces/trace-7520-1.cap >> /mnt/LONTAS/ExpControl/dire15/logs/trace-7518-1.txt