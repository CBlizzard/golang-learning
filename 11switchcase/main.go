package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("its ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1")
	case 2:
		fmt.Println("Dice number is 2")
	case 3:
		fmt.Println("Dice number is 3")
	case 4:
		fmt.Println("chaaaaar")
	case 5:
		fmt.Println("Dice number is 5")
	case 6:
		fmt.Println("Dice number is 6")
	default:
		fmt.Println("try again next time")

	}
}
