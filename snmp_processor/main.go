package main

import "os"
import "log"
import "time"
import "fmt"

func Panic(err error){
	if err != nil {
		log.Println("error: %s",err)
		os.Exit(1)
	}
}

func main(){
	var err error
	err = dbOpen()
	Panic(err)

	// Lists todo experiments
	todo_exp,err := todo_experiments()
	Panic(err)

	for _, entry := range todo_exp {
		expid := entry[0].(int)
		runid := entry[1].(int)
		when_to_process := int64(entry[2].(int))

		if time.Now().UTC().Unix() > when_to_process {
			// converts trace file into a text file using capshow command.
			err = capshow(expid, runid)
			Panic(err)

			// Extracts the required packet timestamps and packetlength from tracefile of two streams(which is a text file after conversion).
			d01_time, d01_length, d10_time, d10_length, err := extract(expid, runid)
			Panic(err)
			// writing influx tables
			intime, err := Influx_Write(d01_time, d01_length, fmt.Sprintf("in_%d_%d", expid, runid))
			Panic(err)

			_, err = Influx_Write(d10_time, d10_length, fmt.Sprintf("out_%d_%d", expid, runid))
			Panic(err)

			time := firsttime(expid, runid) // from snmp logs
			fmt.Println("snmp logs", time.UTC().UnixNano())
			fmt.Println("influx db time", intime.UTC().UnixNano())

			err = sva(expid, runid, intime)
			Panic(err)
		}
	}
}