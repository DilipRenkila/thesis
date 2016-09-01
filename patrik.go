package main

import "os"
import "fmt"
import "encoding/json"
import "bufio"

type InRecord struct {
	Bytes int64 `json:"in_1"`
	SerialID int `json:"serial_id"`
	Unixtime string `json:"unixtime"`
	Uptime float64 `json:"uptime"`
	SwitchManagementIP string `json:"switch_management_ip"`
}
type OutRecord struct {
	Bytes int64 `json:"out_8"`
	SerialID int `json:"serial_id"`
	Unixtime string `json:"unixtime"`
	Uptime float64 `json:"uptime"`
	SwitchManagementIP string `json:"switch_management_ip"`
}

func Decode(r []byte,y string) (x *InRecord,X *OutRecord,err error) {
	if y=="in" {
	    x = new(InRecord)
	    err = json.Unmarshal(r,x)
		fmt.Println(x.SerialID,x.SwitchManagementIP,x.Unixtime,x.Uptime,x.Bytes)
	}
	if y=="out" {
	    X = new(OutRecord)
	    err = json.Unmarshal(r,X)
		fmt.Println(X.SerialID,X.SwitchManagementIP,X.Unixtime,X.Uptime,X.Bytes)
	}
	return
}

func main(){
	args := os.Args
	f, err := os.Open(fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/%s-%s-%s.txt",args[3],args[1],args[2]))
	if err!= nil{
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := []byte(scanner.Text())
		_,_, err:= Decode(input,args[3])
		if err != nil {
			os.Exit(1)
		}
	}

}