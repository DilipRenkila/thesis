package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printCommand(cmd *exec.Cmd) {
  fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
  if err != nil {
    os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  }
}

func printOutput(outs []byte) {
  if len(outs) > 0 {
    fmt.Printf("==> Output: %s\n", string(outs))
  }
}
func main() {
	expid :=7482 ; runid := 1
	//converting tracefile to a text file
	tracefile := fmt.Sprintf("/mnt/LONTAS/traces/trace-%s-%s.cap",expid,runid)
	tracedestiny := fmt.Sprintf("/home/ats/dire15/thesis/logs/trace-%s-%s.txt",expid,runid)
	cmd :=exec.Command("capshow",tracefile,">>",tracedestiny)
	// Create an *exec.Cmd
	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
}
