#!/bin/bash

delay=$1

curl http://10.1.0.119:8080/set_delay_0
curl http://10.1.0.119:8080/delay/$delay
curl http://10.1.0.48:8080/export_delay/$delay
curl http://10.1.0.48:8080/start
# ~/trafficgenerators/udpClient1  -s 192.168.186.221 -n $2 -l $3 -w $4
curl http://10.1.0.48:8080/stop


