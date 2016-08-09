
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/influxdata/influxdb/client"
)


func Influx_Write(d01 [][]string,d10 [][]string) error  {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		return err
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		return err
	}
//    I, err := strconv.ParseInt("1405544147", 10, 64)
	// if err != nil {
        //panic(err)
    //}
    //tm := time.Unix(I, 24767)

	var sampleSize int64
	sampleSize = 1000

	var (
		//shapes     = []string{"circle", "rectangle", "square", "triangle"}
		//colors     = []string{"red", "blue", "green"}
		pts        = make([]client.Point, sampleSize)
	)

	rand.Seed(42)
	var i int64
	for i = 0; i < sampleSize; i++ {
		pts[i] = client.Point{
			Measurement: "shapers1",
		//	Tags: map[string]string{
		//		"color": strconv.Itoa(rand.Intn(len(colors))),
		//		"shape": strconv.Itoa(rand.Intn(len(shapes))),
		//	},
			Fields: map[string]interface{}{
				"value": 123,
			},
			Time:      time.Unix(I,i),
		}
	}

	//fmt.Println(pts)
	bps := client.BatchPoints{
		Points:          pts,
		Database:        "mydb",
		RetentionPolicy: "default",
	}
	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func main(){
	ExampleClient_Write()
}