package main

import ("fmt"
	"strings"
)


func secondmain() {
	info , _ := read_file(fmt.Sprintf("/info/details.txt"))
	infoarray  := strings.Split(info[len(info)-1], " ")
	exp := strings.Split(infoarray[0], ":")
	expid := exp[1]
	fmt.Println(expid)
	return 
}