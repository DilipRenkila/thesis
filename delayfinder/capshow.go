package main

import "os/exec"
import "log"

func main() {
	cmd :=exec.Command("/bin/bash","capshow.sh")
	log.Println(cmd)

}
