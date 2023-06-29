package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println(ptr)  // <nil>

	num := 5
	ptr := &num
	fmt.Println("value that pointer is storign:", ptr)      // something like 0xc0000160b0
	fmt.Println("value that pointer is referencing:", *ptr) // 5

	*ptr = *ptr * 2
	fmt.Println("new value is ", num)

}
