package main

import "os/exec"
import "os"
import "fmt"
import "log"
import "strings"


// Prints the shell command used in os.exec
func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

// Prints the console output to a file.
func printOutput(outs []byte,filename string) {
	if len(outs) > 0 {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666 )
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

// Converts a tracefile to a text file
func capshow(expid int,runid int) error {
	tracefile := fmt.Sprintf("/mnt/LONTAS/traces/trace-%d-%d.cap",expid,runid)
	tracedestiny := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-%d-%d.trace",expid,runid)
	cmd := exec.Command("capshow","-a","-p","20",tracefile)
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	printOutput(output, tracedestiny)
	return nil
}

