package main

import (
	"fmt"
)

func main() {
	// var aobajohsai = []string{"Oikawa", "Iwaizumi", "Matsukawa"}
	// fmt.Println("members are :", aobajohsai)

	// aobajohsai = append(aobajohsai, "Hanamaki", "Yahaba", "Kyotani")
	// fmt.Println(aobajohsai)

	// aobajohsai = append(aobajohsai[1:])
	// fmt.Println(aobajohsai)

	// ==========================

	// highscores := make([]int, 4)
	// highscores[0] = 234
	// highscores[1] = 945
	// highscores[2] = 465
	// highscores[3] = 867
	// highscores = append(highscores, 555, 666, 333)
	// sort.Ints(highscores)
	// fmt.Println(highscores)

	// ==========================
	// remove value from slice based on index
	var teams = []string{"karasuno", "aobajohsai", "shiratorizawa", "nekoma", "fukurodani"}
	fmt.Println(teams)
	index := 2
	teams = append(teams[:index], teams[index+1:]...)
	fmt.Println(teams)
}
