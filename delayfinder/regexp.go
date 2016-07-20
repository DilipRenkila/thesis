package main

import ("fmt"
)


func secondmain() {
	info , _ := read_file(fmt.Sprintf("/info/details.txt"))
	fmt.Println(len(info))
	fmt.Println(info[len(info)-1])
	return 
}