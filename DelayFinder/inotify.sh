#!/bin/bash
inotifywait -m /mnt/LONTAS/traces -e create -e moved_to |
while read action file; do
        echo "The file '$file' appeared in directory '$path' via '$action'"
done
wait