package main

import "os/exec"
import "fmt"

func capshow(expid int,runid int) error {
	cmd := "capshow"
	tracefile := fmt.Sprintf("/mnt/LONTAS/traces/trace-%d-%d.cap",expid,runid)
	tracedest := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-%d-%d.trace",expid,runid)
	args := []string{"-a",tracefile, ">>", tracedest}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		return err
	}
	return nil
}
func main(){
	var err error
	err = capshow(7754,1)
	if err != nil {
		fmt.Println(err)
	}
}