package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"

)

func read_file (path string) ([]string, error) {
	pwd :="/mnt/LONTAS/ExpControl/dire15"
	file, err := os.Open(pwd + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func append_file(lines string) error {
	pwd :="/home/ats/dire15/thesis"
	file, err := os.OpenFile(pwd + "/results.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
    		return err
	}
	defer file.Close()

	if _, err = file.WriteString(lines); err != nil {
    		return err
	}

	return err
}


func main() {


	lines ,err := read_file(fmt.Sprintf("/logs/trace-7541-1.txt"))
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
		fmt.Println(lines_d01[i])
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
			delay := Out - In
			average_delay = average_delay + delay

        	} else {
                	fmt.Println("key not found")
        	}

	}

	x := float64(number_of_packets)
	average_delay = average_delay/x
//	drop := len(lines_d01) - len(lines_d10)
//	delay_in_ms := average_delay*1000
//	fmt.Println(delay_in_ms,drop)


//	err = append_file(fmt.Sprintf("expid:%s runid:%s delay:%s average_delay:%f dropped_packets:%d\n",expid[0],runid[0],delay[0],delay_in_ms,drop))
//	if err != nil {
//		log.Fatalf("append_file: %s",err)
	}

