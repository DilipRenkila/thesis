package main

import "os"
import "fmt"
import "time"

func main() {

	var err error
	err = dbOpen()
	if err != nil {
		fmt.Errorf("error: %s\n", err)
		os.Exit(1)
	}


	// Experiments
	experiments, err := todo_experiments()
	if err != nil {
		fmt.Errorf("error: %s", err)
	}

	for _, entry := range experiments {
		expid := (entry[0].(int))
		runid := entry[1].(int)
		when_to_process := int64(entry[2].(int))

		if  time.Now().Unix() > when_to_process{
			fmt.Println(expid,runid)
			err = capshow(expid,runid)
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}
			filename := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-%d-%d.trace",expid,runid)
			_ , error := read_file(filename)
			if error != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}

		}else {
			time.Sleep( time.Second * 60 )
			fmt.Println(expid,runid)
			err = capshow(expid,runid)
			if err != nil {
				fmt.Errorf("error: %s",err)
			}

		}
	}
}
