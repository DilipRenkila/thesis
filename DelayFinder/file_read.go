package main

import "os"
import "bufio"

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

