package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"encoding/json"
	"github.com/influxdata/influxdb/client"

)
func Influx_Query(query string,query1 string) (int64,int64,error) {

	var size int64 = 0
	var size1 int64= 0
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		return size,size1,err
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		return size,size1,err
	}

	q := client.Query{
		Command:  query,
		Database: "dire15",
	}
	response, err := con.Query(q)
	if err == nil && response.Error() == nil {
		for i := 0;i < len(response.Results[0].Series[0].Values)-1 ; i ++ {
			length,_ := response.Results[0].Series[0].Values[i][1].(json.Number).Int64()
			size = size + length

		}

	} else {
		return size,size1,err
	}
	q1 := client.Query{
		Command:  query1,
		Database: "dire15",
	}
	response1, err := con.Query(q1)
	if err == nil && response1.Error() == nil {
		for i := 0;i < len(response1.Results[0].Series[0].Values)-1 ; i ++ {
			length1,_ := response1.Results[0].Series[0].Values[i][1].(json.Number).Int64()
			size1 = size1 + length1

		}

	} else {
		return size,size1,err
	}
	return size,size1,err
}

func Influx_Write(d []string,d_l []string,tablename string) (time.Time,error)  {
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
	sampleSize = len(d)
	var pts = make([]client.Point, sampleSize)
	var i int
	for i = 0; i < sampleSize  ; i++ {
		timestring := strings.Split(d[i], ".")
		integer_part, _ := strconv.ParseInt(timestring[0], 10, 64)
		decimal_part, _ := strconv.ParseInt(timestring[1], 10, 64)
		value,_ := strconv.ParseInt(d_l[i],10,64)
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
		Database:        "dire15",
		RetentionPolicy: "default",
	}
	_, err = con.Write(bps)
	if err != nil {
		return firsttime,err
	}
	timestring := strings.Split(d[0], ".")
	integer_part, _ := strconv.ParseInt(timestring[0], 10, 64)
	decimal_part, _ := strconv.ParseInt(timestring[1], 10, 64)
	firsttime = time.Unix(integer_part,decimal_part)
	return firsttime,nil
}
