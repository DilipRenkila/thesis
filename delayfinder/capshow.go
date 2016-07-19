package main

import "os/exec"
import "log"

func capshow() {
	cmd :=exec.Command("bash","capshow.sh")
	log.Println(cmd)
	return
}
