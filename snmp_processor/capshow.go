package main

import "os/exec"
import "fmt"

func capshows() error {
	//cmd := "/usr/local/bin/capshow"
	//tracefile := fmt.Sprintf("/mnt/LONTAS/traces/trace-%d-%d.cap",expid,runid)
	//tracedest := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-%d-%d.trace",expid,runid)
	//args := []string{tracefile}
	if err := exec.Command("capshow","/mnt/LONTAS/traces/trace-7754-1.cap").Run(); err != nil {
		return err
	}
	return nil
}
func main(){
	var err error
	err = capshows()
	if err != nil {
		fmt.Println(err)
	}
}