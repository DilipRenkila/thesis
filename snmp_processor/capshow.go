package main

import "os/exec"
import (
	"fmt"
	"os"
)
import "strings"
import "log"
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

func main() {

	//converting tracefile to a text file
	tracefile := fmt.Sprintf("/mnt/LONTAS/traces/trace-7754-1.cap")
	tracedestiny := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-7754-1.trace")
	cmd := exec.Command("capshow", tracefile)
	// Create an *exec.Cmd for executing os commands
	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output, tracedestiny)

}