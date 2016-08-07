package main

import "fmt"
import "regexp"
import "strings"

func extract(filename string) error {

	lines, err := read_file(filename)
	if err != nil {
		return err
	}

	re_1, err := regexp.Compile(`d01`)
	re, err := regexp.Compile(`UDP`)
	reg, err := regexp.Compile(`packets read.`)
	re_2, err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return err
	}

	for _, line := range lines {
		if re_1.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line) == false {
			x := strings.Split(line, ":")
                	y := strings.Split(line, "=")
			z := strings.Split(y[1]," ")
			fmt.Println(x[3],z[0])
		}

		if re_2.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line) == false {
			x := strings.Split(line, ":")
                	y := strings.Split(line, "=")
			z := strings.Split(y[1]," ")
			fmt.Println(x[3],z[0])
		}
	}
	return nil
}