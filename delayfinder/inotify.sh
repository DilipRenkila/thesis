#!/bin/bash
inotifywait -q -m -e close_write /mnt/LONTAS/ExpControl/dire15/status/status |
while read -r filename event; do
  echo "sleep for 150 sec"
  sleep 150
  ./capshow.sh         # or "./$filename"
done