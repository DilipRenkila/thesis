#!/bin/bash
MONITORDIR="/mnt/LONTAS/ExpControl/dire15/logs/"
inotifywait -m -e create --format '%w%f' "${MONITORDIR}"   |
while read file; do
        echo "The file '$file' appeared in directory "
done