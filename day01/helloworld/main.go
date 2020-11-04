package main

import "fmt"

var (
	name string
	age  int
	flag bool
)

func main() {
	name = "111"
	var happy string = "Hello World"
	fmt.Println(happy+name)
	fmt.Scanln()
}
