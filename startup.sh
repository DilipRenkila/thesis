#!/bin/bash
# Argument = -t test -r server -p password -v

usage()
{
cat << EOF
usage: $0 options

OPTIONS:
   -h      Show this message
   -d      your preferred delay on the shaper in microseconds [ required ]
   -s      destination ip address must be in the 192.168.186.0/24 network [ required ]
   -p      destination port number [ optional ]
           defaults to 1500
   -n      number of packets to be sent to destination [ required ]
   -l      minimum packet length in Bytes [ required ]
   -L      maximum packet length in Bytes [ required ]
   -m      packet length distribution (choose e-->exponential, u-->uniform, d-->discrete uniform ) [ optional ]
           defaults to deterministic distribution
   -f      sampling interval in microseconds [ required ]
   -k      keyid [ required ]
   -i      minimum interframe gap, in microseconds [ required ]
   -I      maximum interframe gap, in microseconds [ required ]
   -v      interframe gap distribution ( choose e-->exponential, u-->uniform, d-->discrete uniform ) [ optional ]
           defaults to deterministic distribution

EOF
}
PORT=1500
PACKET_DISTRIBUTION='default'
INTERFRAMEGAP_DISTRIBUTION='default'

while getopts "hd:s:p:n:l:L:m:f:k:i:I:v" OPTION
do
     case $OPTION in
         h)
             usage
             exit 1
             ;;
         d)
             DELAY=$(printf '%dms' "$OPTARG" )
             ;;
         s)
             DESTINATION=$OPTARG
             ;;
         p)
             PORT=$OPTARG
             ;;
         n)
             COUNT=$OPTARG
             ;;
         l)
             MIN_LENGTH=$OPTARG
             ;;
         L)
             MAX_LENGTH=$OPTARG
             ;;
         m)
             PACKET_DISTRIBUTION=$OPTARG
             ;;
         f)
             SAMPLING_INTERVAL=$OPTARG
             ;;
         k)
             KEY_ID=$OPTARG
             ;;
         i)
             MIN_INTERFRAMEGAP=$OPTARG
             ;;
         I)
             MAX_INTERFRAMEGAP=$OPTARG
             ;;
         v)
             INTERFRAMEGAP_DISTRIBUTION=$OPTARG
             ;;
         ?)
             usage
             echo "SUCCESS"
             exit
             ;;
     esac
done


if [[ -z $DELAY ]] || [[ -z $DESTINATION ]] || [[ -z $COUNT ]] || [[ -z $MIN_LENGTH ]] || [[ -z $MAX_LENGTH ]] || [[ -z $SAMPLING_INTERVAL ]] || [[ -z $KEY_ID ]] || [[ -z $MIN_INTERFRAMEGAP ]] || [[ -z $MAX_INTERFRAMEGAP ]]
then
     usage
     echo "SUCCESS"
     exit 1
fi

echo Delay on the shaper is "${DELAY}".
echo This will probe for every "${SAMPLING_INTERVAL}" ms.
epochTime=$(date +%s)
echo expid:$EXPID runid:$RUNID keyid:$KEY_ID delay-on-shaper:$DELAY packets-sent:$COUNT min-packet-length:$MIN_LENGTH max-packet-lenth:$MAX_LENGTH packet-distribution:$PACKET_DISTRIBUTION sampling-interval:$SAMPLING_INTERVAL min-intergramegap:$MIN_INTERFRAMEGAP max-intergramegap:$MAX_INTERFRAMEGAP interframegap-distribution:$INTERFRAMEGAP_DISTRIBUTION destination:$DESTINATION  >> /mnt/LONTAS/ExpControl/dire15/info/details.txt
curl http://10.1.0.119:8080/set_delay_0
curl http://10.1.0.119:8080/delay/$DELAY
sudo service thesis start
/home/ats/trafficgenerators/udpclient -e $EXPID -r $RUNID -k $KEY_ID -s $DESTINATION -p $PORT -n $COUNT -l $MIN_LENGTH -L $MAX_LENGTH -m $PACKET_DISTRIBUTION -w $MIN_INTERFRAMEGAP -W $MAX_INTERFRAMEGAP -v $INTERFRAMEGAP_DISTRIBUTION
sudo service thesis stop
epochTime=$(date +%s)
Time=$((epochTime+300))
mysql -u root -p1 thesis << EOF
    INSERT INTO info (expid,runid,keyid,delay_on_shaper,packets_sent,min_packet_length,max_packet_lenth,packet_distribution,sampling_interval,min_intergramegap,max_intergramegap,interframegap_distribution,destination,status,when_to_process)VALUES ($EXPID,$RUNID,$KEY_ID,'$DELAY',$COUNT,$MIN_LENGTH,$MAX_LENGTH,'$PACKET_DISTRIBUTION',$SAMPLING_INTERVAL,$MIN_INTERFRAMEGAP,$MAX_INTERFRAMEGAP,'$INTERFRAMEGAP_DISTRIBUTION','$DESTINATION',0,$Time);
EOF
echo "SUCCESS"
