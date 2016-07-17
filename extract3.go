package main
import (
	"strings"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"os"
	"bufio"
)
func read_file (path string) ([]string, error) {
	pwd :="/home/ats/dire15/thesis"
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

func main() {

	lines, err := read_file(fmt.Sprintf("/logs/trace-7474-1.txt"))
	if err != nil {
		log.Fatalf("read_file: %s", err)
	}

	re_1, err := regexp.Compile(`d01`)
	re, err := regexp.Compile(`packets read.`)
	re_2, err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	var lines_d01 []string
	var lines_d10 []string
	average_delay := 0.0

	for _, line := range lines {
		if re_1.MatchString(line) == true && re.MatchString(line) == false {
			lines_d01 = append(lines_d01, line)
		}

		if re_2.MatchString(line) == true && re.MatchString(line) == false {
			lines_d10 = append(lines_d10, line)
		}
	}
	m := make(map[string]float64)

	for i, _ := range lines_d01 {
		in := strings.Split(lines_d01[i], ":")
		d01_checksum_string := strings.Split(lines_d01[i], "=")
		In, _ := strconv.ParseFloat(in[3], 64)
		checksum := d01_checksum_string[2]
		m[checksum] = In

	}
	fmt.Println(len(lines_d01),len(lines_d10))
	for j, _ := range lines_d10 {
		out := strings.Split(lines_d10[j], ":")
		d10_checksum_string := strings.Split(lines_d10[j], "=")
		fmt.Println(d10_checksum_string)

		checksum := d10_checksum_string[2]
		Out, _ := strconv.ParseFloat(out[3], 64)
		if _, ok := m[checksum]; ok {
			In := m[checksum]
			delay := Out - In
			average_delay = average_delay + delay
			if delay < 0 {
				fmt.Println(delay,In,lines_d10[j])
			}

        	} else {
                	fmt.Println("key not found")
        	}

	}

	x := float64(len(lines_d10))
	average_delay = average_delay/x
	fmt.Println(average_delay)

}
