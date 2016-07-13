package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
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
	var delay []float64
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
		delay = out[3]-in[3]
		fmt.Println(i,in[3],out[3],delay)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}