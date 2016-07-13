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


func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/logs/trace-7474-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re_1,err := regexp.Compile(`d01`)
	re_2,err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is problem with your regexp.\n")
		return
	}

	var lines_d01 []string
	var lines_d10 []string
	//var average_delay []float64
	average_delay := 0.0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if re_1.MatchString(scanner.Text()) == true  {
		lines_d01 =  append(lines_d01,scanner.Text())
		}

		if re_2.MatchString(scanner.Text()) == true {
		lines_d10 = append(lines_d10,scanner.Text())
		}

	}
	//fmt.Println(lines_d01)
	for i, _ := range lines_d01 {
		//fmt.Println(i, lines_d01[i])
		in := strings.Split(lines_d01[i], ":")
		out := strings.Split(lines_d10[i], ":")
		In, _ := strconv.ParseFloat(in[3], 64)
		Out,_ := strconv.ParseFloat(out[3], 64)
		delay := Out - In
		average_delay = average_delay + delay
	//	fmt.Println(i,in[3],out[3],delay)
	}
	x := len(lines_d01)
	fmt.Println(average_delay / x)
	//fmt.Println(len(lines_d01),len(lines_d10))


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}