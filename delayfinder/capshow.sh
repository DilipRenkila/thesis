#!/bin/bash
input="/mnt/LONTAS/ExpControl/dire15/info/details.txt"
while IFS= read -r var
do
    echo "$var"
done < "$input"
echo "$var"
#capshow /mnt/LONTAS/traces/trace-7520-1.cap >> /mnt/LONTAS/ExpControl/dire15/logs/trace-7518-1.txt