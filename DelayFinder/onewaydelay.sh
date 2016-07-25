#!/bin/bash
input="$(tail -1 /mnt/LONTAS/ExpControl/dire15//info/details.txt)"
set -f
array=(${input// / })
exp=(${array[0]//:/ })
run=(${array[1]//:/ })
expid=${exp[1]}
runid=${run[1]}

/home/ats/consumer-onewaydelay/oneway -a /mnt/LONTAS/traces/trace-${expid}-${runid}.cap >> /mnt/LONTAS/ExpControl/dire15/logs/trace-${expid}-${runid}.txt