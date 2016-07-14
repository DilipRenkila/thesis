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
	pwd , _ := os.Getwd()
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

	file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_WRONLY, 0600)
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

	lines ,err := read_file("/logs/trace-7474-1.txt")
	if err != nil {
		log.Fatalf("read_file: %s",err)
	}
	delay,err := read_file("/delay.txt")
	if err != nil {
		log.Fatalf("read_file: %s",err)
	}

	expid,err := read_file("/expid.txt")
	if err != nil {
		log.Fatalf("read_file: %s",err)
	}

	runid,err := read_file("/runid.txt")
	if err != nil {
		log.Fatalf("read_file: %s",err)
	}

	re_1,err := regexp.Compile(`d01`)
	re_2,err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is problem with your regexp.\n")
		return
	}

	var lines_d01 []string
	var lines_d10 []string
	average_delay := 0.0

	for _,line := range lines {
		if re_1.MatchString(line) == true  {
		lines_d01 =  append(lines_d01,line)
		}

		if re_2.MatchString(line) == true {
		lines_d10 = append(lines_d10,line)
		}
	}

	for i, _ := range lines_d01 {
		in := strings.Split(lines_d01[i], ":")
		out := strings.Split(lines_d10[i], ":")
		In, _ := strconv.ParseFloat(in[3], 64)
		Out,_ := strconv.ParseFloat(out[3], 64)
		delay := Out - In
		average_delay = average_delay + delay
	}
	x := float64(len(lines_d01))
	fmt.Println(average_delay / x)
	err = append_file(fmt.Sprintf("expid:%s runid:%s delay:%s average_delay:%s",expid,runid,delay,average_delay))
	if err != nil {
		log.Fatalf("append_file: %s",err)
	}

}