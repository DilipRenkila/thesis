package main

import "os/exec"
import "os"
import "fmt"
import "log"
import "strings"


func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

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

func capshow(expid int,runid int) error {

	//converting tracefile to a text file
	tracefile := fmt.Sprintf("-",expid,runid)
	tracedestiny := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-%d-%d.trace",expid,runid)
	cmd := exec.Command("capshow","-a",tracefile)
	// Create an *exec.Cmd for executing os commands
	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	printOutput(output, tracedestiny)
	return nil

}