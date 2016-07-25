#!/bin/bash
inotifywait -q -m -e close_write /mnt/LONTAS/ExpControl/dire15/status/status |
while read -r filename event; do
  echo "sleeping for 150 sec"
  sleep 150
  bash onewaydelay.sh
done