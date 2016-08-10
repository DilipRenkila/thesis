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

			err = capshow(expid,runid)
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}
			filename := fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/trace-%d-%d.trace",expid,runid)
			d01,d10,err := extract(filename)
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}
			in_table := fmt.Sprintf("in_%d_%d",expid,runid)
			out_table := fmt.Sprintf("out_%d_%d",expid,runid)
			intime,err := Influx_Write(d01,in_table)
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}

			_,err = Influx_Write(d10,out_table)
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}
			sva(expid,runid,intime)
			//ExampleClient_Query()

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
