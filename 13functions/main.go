package main

import "fmt"

func main() {
	result := add2(1, 2)
	fmt.Println(result)

	result2, message := addN(1, 2, 3, 4, 5)
	fmt.Println(message, result2)
}

func add2(num1 int, num2 int) int {
	return num1 + num2
}

func addN(values ...int) (int, string) {
	result := 0
	for _, val := range values {
		result += val
	}
	return result, "heres the result :"
}
