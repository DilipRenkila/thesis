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

	lines, err := read_file(fmt.Sprintf("/logs/trace-7485-1.txt"))
	if err != nil {
		log.Fatalf("read_file: %s", err)
	}

	re_1, err := regexp.Compile(`d01`)
	re_2, err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	var lines_d01 []string
	var lines_d10 []string

	for _, line := range lines {
		if re_1.MatchString(line) == true {
			lines_d01 = append(lines_d01, line)
		}

		if re_2.MatchString(line) == true {
			lines_d10 = append(lines_d10, line)
		}
	}

	for i, _ := range lines_d01 {
		in := strings.Split(lines_d01[i], ":")
		d01_checksum_string := strings.Split(lines_d01[i], "=")
		for j, _ := range lines_d10 {
			d10_checksum_string := strings.Split(lines_d01[j], "=")
			out := strings.Split(lines_d10[j], ":")
			if d01_checksum_string[3] == d10_checksum_string[3] {
				In, _ := strconv.ParseFloat(in[3], 64)
				Out, _ := strconv.ParseFloat(out[3], 64)
				fmt.Println(d01_checksum_string[3], Out - In)

			}
		}
	}

}
