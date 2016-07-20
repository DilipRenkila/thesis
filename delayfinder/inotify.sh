#!/bin/bash
inotifywait -q -m -e close_write /mnt/LONTAS/ExpControl/dire15/status/status |
while read -r filename event; do
  echo "sleep for 15 sec"
  sleep 15
  bash capshow.sh         # or "./$filename"
done