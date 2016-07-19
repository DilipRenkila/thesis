package main

import "os/exec"
import "log"

func capshow() {
	cmd :=exec.Command("/bin/bash","capshow.sh")
	log.Println(cmd)
	return
}
