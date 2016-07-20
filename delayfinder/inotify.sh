#!/bin/bash
inotifywait -q -m -e close_write /mnt/LONTAS/ExpControl/dire15//info/details.txt |
while read -r filename event; do
  sleep 120
  ./capshow.sh         # or "./$filename"
done