package main

import ("fmt"
	"strings"
	"log"
	"regexp"
	"strconv"

)
type info struct {
  expid,runid,interframegap,destination,samplinginterval,packetlength,packets,delayonshaper,packets string
}

func main() {
	info , _ := read_file(fmt.Sprintf("/info/details.txt"))
	infoarray  := strings.Split(info[len(info)-1], " ")

	exp := strings.Split(infoarray[0], ":")
	run := strings.Split(infoarray[1], ":")
	del := strings.Split(infoarray[3], ":")
	pack := strings.Split(infoarray[4], ":")
	packlen := strings.Split(infoarray[5], ":")
	sampint := strings.Split(infoarray[6], ":")
	dest := strings.Split(infoarray[7], ":")
	interframe := strings.Split(infoarray[8], ":")
	Info := info{exp[1],run[1],interframe[1],dest[1],sampint[1],packlen[1],del[1],pack[1]}

	lines ,err := read_file(fmt.Sprintf("/logs/trace-%s-%s.txt",Info.expid,Info.runid))
	if err != nil {
		log.Fatalf("read_file: %s", err)
	}

	re_1, err := regexp.Compile(`mp10165_d01`)
	re_2, err := regexp.Compile(`mp10165_d10`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
	}

	average_delay := 0.0
	number_of_packets := 0

	for _, line := range lines {
		if re_1.MatchString(line) == true && re_2.MatchString(line) == true {

			in := strings.Split(line, ";")

			delay, _ := strconv.ParseFloat(in[5], 64)
			number_of_packets = number_of_packets + 1
			average_delay = average_delay + delay
		}
	}

	x := float64(number_of_packets)
	average_delay = average_delay/x
	delay_in_ms := average_delay*1000
	err = append_file(fmt.Sprintf("expid:%s runid:%s delay-on-shaper:%s average_delay:%f packets_sent:%s packets_received:%d packet_length:%s sampling_interval_in_sec:%s destination:%s interframegap:%s\n",Info.expid,Info.runid,Info.delayonshaper,delay_in_ms,Info.packets,number_of_packets,Info.packetlength,Info.samplinginterval,Info.destination,Info.interframegap))
	if err != nil {
		log.Fatalf("append_file: %s",err)
	}

}