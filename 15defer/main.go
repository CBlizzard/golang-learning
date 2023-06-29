package main

import "fmt"

// defer => LIFO
func main() {
	fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Four")
	reversePrint()
}

func reversePrint() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// output
// One
// Four
// 4
// 3
// 2
// 1
// 0
// Three
// Two
