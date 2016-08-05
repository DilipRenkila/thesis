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
		var current_line []string
		current_line = append(current_line,scanner.Text())
		infoarray  := strings.Split(current_line[0], " ")
		fmt.Println(infoarray[4])
//		status := strings.Split(infoarray[9], ":")
//		if status[1] == "pending" {
			lines = append(lines, scanner.Text())
//		}
	}
	return lines, scanner.Err()
}

func main() {
	_,err := read_file()

	fmt.Println(err)
}
