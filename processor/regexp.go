package main

import "fmt"
import "regexp"
import "strings"

func extract(filename string) ([]string,[]string,[]string,[]string,error) {
	var d01_time,d10_time []string
	var d01_length,d10_length []string
	lines, err := read_file(filename)
	if err != nil {
		return d01_time,d01_length,d10_time,d10_length,nil
	}

	re_1, err := regexp.Compile(`d01`)
	re, err := regexp.Compile(`UDP`)
	reg, err := regexp.Compile(`packets read.`)
	re_2, err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
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