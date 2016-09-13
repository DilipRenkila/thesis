package main

import "log"
import "fmt"
import "regexp"
import "strings"

// Extracts the required packets and their respective timestamps and packetlength from trace file(which is a text file after conversion).
func extract(expid int,runid int) ([]string,[]string,[]string,[]string,error) {
	filename := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-%d-%d.trace",expid,runid)
	var d01_time,d10_time []string
	var d01_length,d10_length []string
	lines, err := read_file(filename)
	if err != nil {
		log.Println("There is a problem while reading file:%s",filename)
		return d01_time,d01_length,d10_time,d10_length,nil
	}
	re_1, err := regexp.Compile(`d01`)
	re, err := regexp.Compile(`UDP`)
	reg, err := regexp.Compile(`packets read.`)
	re_2, err := regexp.Compile(`d10`)
	if err != nil {
		log.Println("There is a problem with your regexp")
		return d01_time,d01_length,d10_time,d10_length,nil
	}

	for _, line := range lines {
		if re_1.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line) == false {
			x := strings.Split(line, ":")
                	y := strings.Split(line, "=")
			z := strings.Split(y[1]," ")
			d01_time = append(d01_time,x[3])
			d01_length=append(d01_length,z[0])
		}
		if re_2.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line) == false {
			x := strings.Split(line, ":")
                	y := strings.Split(line, "=")
			z := strings.Split(y[1]," ")
			d10_time = append(d10_time,x[3])
			d10_length=append(d10_length,z[0])
		}
	}
	return d01_time,d01_length,d10_time,d10_length,nil
}