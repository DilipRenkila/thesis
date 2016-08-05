package main

import "os"
import "bufio"
import "strings"
import "fmt"

func read_file () ([]string, error) {
	pwd :="/mnt/LONTAS/ExpControl/dire15"
	path := "/info/details.txt"
	file, err := os.Open(pwd + path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		infoarray  := strings.Split(scanner.Text(), " ")
		status := strings.Split(infoarray[9], ":")
		if status[1] == "pending" {
			lines = append(lines, scanner.Text())
		}
	}
	return lines, scanner.Err()
}

func main() {
	x,_ := read_file()
	fmt.Println(x)
}
