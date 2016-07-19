package main

import "os/exec"
import "log"

func main() {
	cmd :=exec.Command("/home/ats/dire15/thesis/delayfinder/capshow.sh")
	log.Println(cmd)

}
