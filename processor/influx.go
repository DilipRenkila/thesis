package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/influxdata/influxdb/client"
)


func Influx_Write(d [][]string) error  {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		return err
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		return err
	}

	var sampleSize int
	sampleSize = len(d[0])

	var pts = make([]client.Point, sampleSize)
	var i int
	fmt.Println(len(d[0]))
	for i = 0; i < sampleSize  ; i++ {
		timestring := strings.Split(d[0][i], ".")
		fmt.Println(i)
		integer_part, _ := strconv.ParseInt(timestring[0], 10, 64)
		decimal_part, _ := strconv.ParseInt(timestring[1], 10, 64)
		value,_ := strconv.ParseInt(d[1][i],10,64)
		unixtime := time.Unix(integer_part,decimal_part)
		pts[i] = client.Point{
			Measurement: "sample",
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
		return err
	}
	return nil
}
