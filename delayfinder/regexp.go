package main

import ("fmt"
)


func secondmain() {
	info , _ := read_file(fmt.Sprintf("/info/details.txt"))
	fmt.Println(info[len(info)])
	return 
}