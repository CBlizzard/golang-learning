package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	msg := "hey hey hey"
	fmt.Println(msg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("you entered THIS :", input)
	fmt.Printf("type of input is %T", input)
}
