#!/bin/bash
inotifywait -m /mnt/LONTAS/traces -e create -e moved_to |
while read /mnt/LONTAS/traces action file; do
        echo "The file '$file' appeared in directory '$path' via '$action'"
done
wait