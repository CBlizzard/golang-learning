package main

import "fmt"

func main() {
	teams := []string{"karasuno", "aoba johsai", "shiratorizawa", "neko", "fukurodani", "inarizaki", "dateko"}

	// for i := 0; i < len(teams); i++ {
	// 	fmt.Println(teams[i])
	// }

	for i, team := range teams {
		fmt.Println(i, team)
	}

}
