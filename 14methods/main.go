package main

import "fmt"

func main() {
	kakashi := User{"kakashi", "kakashi@hokage.nin", 6}

	kakashi.WhichHokage()
}

type User struct {
	Name        string
	Email       string
	HokageOrder int
}

func (u User) WhichHokage() {
	fmt.Printf("%v is hokage number %v", u.Name, u.HokageOrder)
}
