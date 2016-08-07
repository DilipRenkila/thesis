package main
import "encoding/json"
import "os"
import "fmt"
import "bufio"
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


func main() {
	var Bytes_in []int64
	var uptime []float64
	var bitrate []float64
	f, _ := os.Open("/mnt/LONTAS/ExpControl/dire15/logs/in-7742-1.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := []byte(scanner.Text())
		x, _ := Decode(input)
		Bytes_in = append(Bytes_in,x.In1)
		uptime = append(uptime,x.Uptime)
	}

	for i := len(uptime)-1; i > 0; i-- {
		fmt.Println(float64(Bytes_in[i]-Bytes_in[i-1]))
		bitrate=append(bitrate,(float64(Bytes_in[i]-Bytes_in[i-1])/uptime[i]-uptime[i-1]))
	}
	fmt.Println(bitrate)
}
