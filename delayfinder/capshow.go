package main

import "os/exec"
import "log"

func main() {
	cmd :=exec.Command("capshow.sh")
	log.Println(cmd)

}
