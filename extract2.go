package main

import (
    "log"
//    "fmt"
    "golang.org/x/exp/inotify"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
	"os/exec"
	"time"
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

func append_file(lines string) error {
	pwd :="/home/ats/dire15/thesis"
	file, err := os.OpenFile(pwd + "/results.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
    		return err
	}
	defer file.Close()

	if _, err = file.WriteString(lines); err != nil {
    		return err
	}

	return err
}
func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
  if err != nil {
    os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  }
}

func printOutput(outs []byte,filename string) {
  if len(outs) > 0 {
 //   fmt.Printf("==> Output: %s\n", string(outs))
	//  pwd := "/home/ats/dire15/thesis/logs"
	  file, err := os.OpenFile(
         filename,
        os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
        0666,
    )
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Write bytes to file
    bytesWritten, err := file.Write(outs)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Wrote %d bytes to %s .\n", bytesWritten,filename)

  }
}

func extract() {


	delay,err := read_file("/delay.txt")
	if err != nil {
		log.Fatalf("read_file: %s",err)
	}

	expid,err := read_file("/expid/expid.txt")
	if err != nil {
		log.Fatalf("read_file: %s",err)
	}

	runid,err := read_file("/runid.txt")
	if err != nil {
		log.Fatalf("read_file: %s",err)
	}

	//converting tracefile to a text file
	tracefile := fmt.Sprintf("/mnt/LONTAS/traces/trace-%s-%s.cap",expid[0],runid[0])
	tracedestiny := fmt.Sprintf("/home/ats/dire15/thesis/logs/trace-%s-%s.txt",expid[0],runid[0])
	cmd :=exec.Command("capshow",tracefile)
	// Create an *exec.Cmd for executing os commands
	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output,tracedestiny)


	lines ,err := read_file(fmt.Sprintf("/logs/trace-%s-%s.txt",expid[0],runid[0]))
	if err != nil {
		log.Fatalf("read_file: %s", err)
	}

	re_1, err := regexp.Compile(`d01`)
	re, err := regexp.Compile(`UDP`)
	reg, err := regexp.Compile(`packets read.`)
	re_2, err := regexp.Compile(`d10`)

	if err != nil {
		fmt.Printf("There is a problem with your regexp.\n")
		return
	}

	var lines_d01 []string
	var lines_d10 []string
	average_delay := 0.0
	number_of_packets := 0

	for _, line := range lines {
		if re_1.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line )== false{
			lines_d01 = append(lines_d01, line)
		}

		if re_2.MatchString(line) == true && re.MatchString(line) == true && reg.MatchString(line)== false{
			lines_d10 = append(lines_d10, line)
		}
	}
	m := make(map[string]float64)

	for i, _ := range lines_d01 {
		in := strings.Split(lines_d01[i], ":")
		d01_checksum_string := strings.Split(lines_d01[i], "=")
		In, _ := strconv.ParseFloat(in[3], 64)
		fmt.Println(d01_checksum_string)
		checksum := d01_checksum_string[2]
		m[checksum] = In

	}
	fmt.Println(len(lines_d01),len(lines_d10))
	for j, _ := range lines_d10 {
		out := strings.Split(lines_d10[j], ":")
		d10_checksum_string := strings.Split(lines_d10[j], "=")
		//fmt.Println(d10_checksum_string)

		checksum := d10_checksum_string[2]
		Out, _ := strconv.ParseFloat(out[3], 64)
		if _, ok := m[checksum]; ok {
			number_of_packets = number_of_packets + 1
			In := m[checksum]
			delay := In - Out
			average_delay = average_delay + delay

        	} else {
                	fmt.Println("key not found")
        	}

	}

	x := float64(number_of_packets)
	average_delay = average_delay/x
	drop := len(lines_d01) - len(lines_d10)
	delay_in_ms := average_delay*1000
	fmt.Println(delay_in_ms,drop)


	err = append_file(fmt.Sprintf("expid:%s runid:%s delay:%s average_delay:%f dropped_packets:%d\n",expid[0],runid[0],delay[0],delay_in_ms,drop))
	if err != nil {
		log.Fatalf("append_file: %s",err)
	}

}

func main() {
    watcher, err := inotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan bool)

    // Process events
    go func() {
        for {
            select {
            case ev := <-watcher.Event:
		    if ev.Mask&inotify.IN_MODIFY != inotify.IN_MODIFY {
			    continue
		    }
		    log.Println("Reloading configuration")
                    log.Println(ev.Name)
		    //wait_time,err := read_file("/expid/expid.txt")
			//if err != nil {
			//	log.Fatalf("read_file: %s",err)
			//}

	    		time.Sleep(time.Second*300)
		    extract()

            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch("/home/ats/dire15/thesis/expid")
    if err != nil {
        log.Fatal(err)
    }

    // Hang so program doesn't exit
    <-done
    /* ... do stuff ... */
    watcher.Close()
}