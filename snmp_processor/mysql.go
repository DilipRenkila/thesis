package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"time"
	"os"
)

// Global variables
var db *sql.DB

func dbOpen() error {
	var err error

	db, err = sql.Open("mysql", "root:1@(127.0.0.1:3306)/thesis?charset=utf8")
	if err != nil {
		return err
	}
	return nil
}

func main() {

	var err error
	err = dbOpen()
	if err != nil {
		fmt.Errorf("error: %s\n", err)
		os.Exit(1)
	}


	// Experiments
	experiments, err := todo_experiments()
	if err != nil {
		fmt.Errorf("error: %s", err)
	}

	for _, entry := range experiments {
		expid := (entry[0].(int))
		runid := entry[1].(int)
		when_to_process := int64(entry[2].(int))

		if  time.Now().Unix() > when_to_process{
			fmt.Println(expid,runid)
		}else {
			duration := when_to_process - (time.Now().Unix())
			timeDuration, _ := time.ParseDuration(fmt.Sprintf("%ds", duration))
			time.Sleep( time.Second* timeDuration )
			fmt.Println("slept for %s seconds", duration)
			fmt.Println(expid,runid)

		}
	}
}

func todo_experiments() ([][]interface{}, error) {
	q := fmt.Sprintf("SELECT expid, runid, when_to_process FROM info WHERE status=0;")
	var expid int
	var runid int
	var when_to_process int
	outfmt := []interface{}{expid, runid, when_to_process}
	result, err := dbQueryScan(db, q, nil, outfmt)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func dbDoQueryScan(db *sql.DB, q string, args []interface{}, outargs []interface{}) ([][]interface{}, error) {
	rows, err := db.Query(q, args...)
	if err != nil {
		return [][]interface{}{}, err
	}
	defer rows.Close()
	result := [][]interface{}{}
	for rows.Next() {
		ptrargs := make([]interface{}, len(outargs))
		for i := range outargs {
			switch t := outargs[i].(type) {
			case string:
				str := ""
				ptrargs[i] = &str
			case int:
				integer := 0
				ptrargs[i] = &integer
			default:
				return [][]interface{}{}, fmt.Errorf("Bad interface type: %s\n", t)
			}
		}
		err = rows.Scan(ptrargs...)
		if err != nil {
			return [][]interface{}{}, err
		}
		newargs := make([]interface{}, len(outargs))
		for i := range ptrargs {
			switch t := outargs[i].(type) {
			case string:
				newargs[i] = *ptrargs[i].(*string)
			case int:
				newargs[i] = *ptrargs[i].(*int)
			default:
				return [][]interface{}{}, fmt.Errorf("Bad interface type: %s\n", t)
			}
		}
		result = append(result, newargs)
	}
	err = rows.Err()
	if err != nil {
		return [][]interface{}{}, err
	}
	return result, nil
}

func dbQueryScan(db *sql.DB, q string, inargs []interface{}, outfmt []interface{}) ([][]interface{}, error) {
	for {
		result, err := dbDoQueryScan(db, q, inargs, outfmt)
		if err == nil {
			return result, nil
		}
		time.Sleep(1 * time.Second)
	}
}
