package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:1@(127.0.0.1:3306)/thesis?charset=utf8")
	checkErr(err)
	// query
	rows, err := db.Query("SELECT * FROM info where status=0")
	checkErr(err)

	for rows.Next() {
		var expid,runid,keyid,packets_sent,min_packet_length,max_packet_lenth int
		var sampling_interval,min_intergramegap,max_intergramegap,status,when_to_process int
		var delay_on_shaper,packet_distribution,interframegap_distribution,destination string
		err = rows.Scan(&expid,&runid,&keyid,&delay_on_shaper,&packets_sent,&min_packet_length,&max_packet_lenth,&packet_distribution,&sampling_interval,&min_intergramegap,&max_intergramegap,&interframegap_distribution,&destination,&status,&when_to_process)
		checkErr(err)
		unixtime := int64(time.Now().Unix())
		if (unixtime > when_to_process ){
			fmt.Println("ok, you can process")
		}

	}

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}