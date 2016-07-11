package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/logs/trace-7474-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines_d01 []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lines_d01 =  append(lines_d01,scanner.Text())
	}
	fmt.Println(lines_d01)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}