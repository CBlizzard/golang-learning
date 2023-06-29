package main

import "fmt"

func main() {
	kakashi := User{"kakashi", "kakashi@hokage.nin", 30, true}

	fmt.Println(kakashi)
	fmt.Println(kakashi.Name)
	fmt.Printf("full details %+v", kakashi)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Hokage bool
}
