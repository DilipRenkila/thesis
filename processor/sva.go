package main
import "encoding/json"
import "os"
import "fmt"
import "bufio"
import "time"
//import "reflect"
type InRecord struct {
	In1 int64 `json:"in_1"`
	SerialID int `json:"serial_id"`
	Unixtime string `json:"unixtime"`
	Uptime float64 `json:"uptime"`
	SwitchManagementIP string `json:"switch_management_ip"`
}

type OutRecord struct {
	In1 int64 `json:"out_8"`
	SerialID int `json:"serial_id"`
	Unixtime string `json:"unixtime"`
	Uptime float64 `json:"uptime"`
	SwitchManagementIP string `json:"switch_management_ip"`
}

func InDecode(r []byte) (x *InRecord, err error) {
    x = new(InRecord)
    err = json.Unmarshal(r,x)
    return
}

func OutDecode(r []byte) (x *OutRecord, err error) {
    x = new(OutRecord)
    err = json.Unmarshal(r,x)
    return
}

func timemachine(intime time.Time, interval float64) time.Time {
	ttt := fmt.Sprintf("%fs",interval)
	dur, _ := time.ParseDuration(ttt)
	outtime := intime.Add(dur)
	return outtime
}

func sva(expid int,runid int,intime time.Time) error {
	var Bytes_in []int64
	var Bytes_out []int64
	var X []float64

	var in_uptime []float64
	var out_uptime []float64

	var inbitrate []float64
	var outbitrate []float64

	f, err := os.Open(fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/in-%d-%d.txt",expid,runid))
	if err!= nil{
		return err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := []byte(scanner.Text())
		x, _ := InDecode(input)
		Bytes_in = append(Bytes_in,x.In1)
		in_uptime = append(in_uptime,x.Uptime)
	}

	f, err = os.Open(fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/out-%d-%d.txt",expid,runid))
	if err!= nil{
		return err
	}
	scanner = bufio.NewScanner(f)
	for scanner.Scan() {
		input := []byte(scanner.Text())
		x, _ := OutDecode(input)
		Bytes_out = append(Bytes_out,x.In1)
		out_uptime = append(out_uptime,x.Uptime)
	}

	for i := 0; i < len(in_uptime)-1; i++ {
		inbitrate=append(inbitrate,(float64(Bytes_in[i+1]-Bytes_in[i])/in_uptime[i+1]-in_uptime[i]))
		outbitrate=append(outbitrate,(float64(Bytes_out[i+1]-Bytes_out[i])/out_uptime[i+1]-out_uptime[i]))
		interval := in_uptime[i+1]-in_uptime[i]
		outtime := timemachine(intime,interval)
		str:=fmt.Sprintf("select * from in_%d_%d where time > %v and time < %v",expid,runid,intime.UnixNano(),outtime.UnixNano())
		str1:=fmt.Sprintf("select * from out_%d_%d where time > %v and time < %v",expid,runid,intime.UnixNano(),outtime.UnixNano())
		size,size1,err := Influx_Query(str,str1)
		if err!=nil{
			return err
		}
		if i==0 {
			X=append(X,float64(size-size1))
		} else {
			X=append(X,X[i-1]+float64(size-size1))
		}
		fmt.Println(fmt.Sprintf("interval:%d, range:%v-%v, Bytes_src:%d, Bytes_dest:%d, inbitrate:%f, outbitrate:%f, aditional delay: %f ", i+1,intime.UnixNano(),outtime.UnixNano(),size,size1,inbitrate[i],outbitrate[i],X[i]/inbitrate[i]))

		intime = outtime
		
	}


	return nil
}
