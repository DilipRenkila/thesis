package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"log"
	"reflect"
	"github.com/influxdata/influxdb/client"
)
func ExampleClient_Query() {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		log.Fatal(err)
	}

	q := client.Query{
		Command:  "select  * from in_7754_1",
		Database: "thesis",
	}
	if response, err := con.Query(q); err == nil && response.Error() == nil {
		fmt.Println(reflect.TypeOf(response.Results))
	}
}

func Influx_Write(d [][]string,tablename string) (time.Time,error)  {
	var firsttime time.Time
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		return firsttime,err
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		return firsttime,err
	}

	var sampleSize int
	sampleSize = len(d[0])

	var pts = make([]client.Point, sampleSize)
	var i int
	fmt.Println(len(d[0]),len(d[1]))
	for i = 0; i < sampleSize  ; i++ {
		timestring := strings.Split(d[0][i], ".")
		integer_part, _ := strconv.ParseInt(timestring[0], 10, 64)
		decimal_part, _ := strconv.ParseInt(timestring[1], 10, 64)
		value,_ := strconv.ParseInt(d[1][i],10,64)
		unixtime := time.Unix(integer_part,decimal_part)
		pts[i] = client.Point{
			Measurement: tablename,
			Fields: map[string]interface{}{
				"value": value,
			},
			Time:  unixtime,
		}
	}

	//fmt.Println(pts)
	bps := client.BatchPoints{
		Points:          pts,
		Database:        "thesis",
		RetentionPolicy: "default",
	}
	_, err = con.Write(bps)
	if err != nil {
		return firsttime,err
	}
	timestring := strings.Split(d[0][0], ".")
	integer_part, _ := strconv.ParseInt(timestring[0], 10, 64)
	decimal_part, _ := strconv.ParseInt(timestring[1], 10, 64)
	firsttime = time.Unix(integer_part,decimal_part)
	return firsttime,nil
}
