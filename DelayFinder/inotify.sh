#!/bin/bash
MONITORDIR="/mnt/LONTAS/traces/"
inotifywait -m -e create --format '%w%f' "${MONITORDIR}"   |
while read file; do
        echo "The file '$file' appeared in directory "
done