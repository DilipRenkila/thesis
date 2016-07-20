package main

import ("fmt"
	"strings"
	"log"
	"regexp"
	"strconv"
	
)


func secondmain() {
	info , _ := read_file(fmt.Sprintf("/info/details.txt"))
	infoarray  := strings.Split(info[len(info)-1], " ")
	del := strings.Split(infoarray[2], ":")
	exp := strings.Split(infoarray[0], ":")
	run := strings.Split(infoarray[1], ":")
	pack := strings.Split(infoarray[3], ":")
	packlen := strings.Split(infoarray[4], ":")
	sampint := strings.Split(infoarray[5], ":")
	dest := strings.Split(infoarray[6], ":")
	interframe := strings.Split(infoarray[7], ":")
	delayonshaper := del[1]
	packets := pack[1]
	packetlength := packlen[1]
	samplinginterval := sampint[1]
	destination := dest[1]
	interframegap := interframe[1]
	expid := exp[1]
	runid := run[1]
	lines ,err := read_file(fmt.Sprintf("/logs/trace-%s-%s.txt",expid,runid))
	if err != nil {
		log.Fatalf("read_file: %s", err)
	}

	re_1, err := regexp.Compile(`d01`)
	re, err := regexp.Compile(`UDP`)
	reg, err := regexp.Compile(`packets read.`)
	re_2, err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	var lines_d01 []string
	var lines_d10 []string
	average_delay := 0.0
	number_of_packets := 0

	for _, line := range lines {
		if re_1.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line )== false{
			lines_d01 = append(lines_d01, line)
		}

		if re_2.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line )== false {
			lines_d10 = append(lines_d10, line)
		}
	}
	m := make(map[string]float64)

	for i, _ := range lines_d01 {
		in := strings.Split(lines_d01[i], ":")
		d01_checksum_string := strings.Split(lines_d01[i], "=")
		In, _ := strconv.ParseFloat(in[3], 64)
		checksum := d01_checksum_string[2]
		m[checksum] = In
	}
	for j, _ := range lines_d10 {
		out := strings.Split(lines_d10[j], ":")
		d10_checksum_string := strings.Split(lines_d10[j], "=")
		checksum := d10_checksum_string[2]
		Out, _ := strconv.ParseFloat(out[3], 64)
		if _, ok := m[checksum]; ok {
			number_of_packets = number_of_packets + 1
			In := m[checksum]
			delay := In - Out
			average_delay = average_delay + delay

        	} else {
                	fmt.Println("key not found")
        	}

	}

	x := float64(number_of_packets)
	average_delay = average_delay/x
	delay_in_ms := average_delay*1000
	fmt.Println(delay_in_ms)


	err = append_file(fmt.Sprintf("expid:%s runid:%s delay-on-shaper:%s average_delay:%f packets_sent:%s packets_on_d01:%d packets_packet_d10:%d packet_length:%s sampling_interval_in_sec:%s destination:%s interframegap:%s\n",expid,runid,delayonshaper,delay_in_ms,packets,len(lines_d01),len(lines_d10),packetlength,samplinginterval,destination,interframegap))
	if err != nil {
		log.Fatalf("append_file: %s",err)
	}

	return 
}