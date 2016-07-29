#!/bin/bash
inotifywait -m /mnt/LONTAS/ExpControl/dire15/logs/ -e create  |
while read file; do
        echo "The file '$file' appeared in directory "
done
wait