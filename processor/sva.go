package main
import "encoding/json"
import "os"
import "fmt"
import "bufio"
import "time"
type Record struct {
	In1 int64 `json:"in_1"`
	SerialID int `json:"serial_id"`
	Unixtime string `json:"unixtime"`
	Uptime float64 `json:"uptime"`
	SwitchManagementIP string `json:"switch_management_ip"`
}


func Decode(r []byte) (x *Record, err error) {
    x = new(Record)
    err = json.Unmarshal(r,x)
    return
}

func timemachine(intime time.Time, interval float64) time.Time {
	ttt := fmt.Sprintf("%fs",interval)
	dur, _ := time.ParseDuration(ttt)
	outtime := intime.Add(dur)
	return outtime
}

func sva(expid int,runid int,intime time.Time) {
	var Bytes_in []int64
	var uptime []float64
	var bitrate []float64
	var interval []float64 = 0.0
	var intervals []float64
	f, _ := os.Open(fmt.Sprintf("/mnt/LONTAS/ExpControl/dire15/logs/in-%d-%d.txt",expid,runid))
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := []byte(scanner.Text())
		x, _ := Decode(input)
		Bytes_in = append(Bytes_in,x.In1)
		uptime = append(uptime,x.Uptime)
	}

	for i := 0; i < len(uptime)-1; i++ {
		bitrate=append(bitrate,(float64(Bytes_in[i+1]-Bytes_in[i])/uptime[i+1]-uptime[i]))
		interval = interval + uptime[i+1]-uptime[i]
		intervals = append(intervals,timemachine(intime,interval))
	}
	fmt.Println(intervals)
}
