package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)


func main() {
	file, err := os.Open("~/dire15/thesis/logs/trace-7474-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}