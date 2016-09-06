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


	// searching for experiments todo
	experiments, err := todo_experiments()
	if err != nil {
		fmt.Errorf("error: %s", err)
		os.Exit(1)
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
			d01_time,d01_length,d10_time,d10_length,err := extract(filename)
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}
			in_table := fmt.Sprintf("in_%d_%d",expid,runid)
			out_table := fmt.Sprintf("out_%d_%d",expid,runid)
			intime,err := Influx_Write(d01_time,d01_length,in_table) // from influx tables
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}

			_,err = Influx_Write(d10_time,d10_length,out_table)
			if err != nil {
				fmt.Errorf("error: %s", err)
				os.Exit(1)
			}
			time := firsttime(expid,runid) // from snmp logs
			fmt.Println("snmp logs",time.UnixNano())
			fmt.Println("influx db time",intime.UnixNano())

			err=sva(expid,runid,intime)
			if err != nil {
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
