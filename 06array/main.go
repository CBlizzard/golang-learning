package main

import "fmt"

func main() {
	var karasuno [5]string

	karasuno[0] = "Hinata"
	karasuno[1] = "Kageyama"
	karasuno[2] = "Tsukishima"
	karasuno[4] = "Tanaka"

	fmt.Println("members are :", karasuno)

	var nekoma = [5]string{"Kuroo", "Kenma", "Lev", "Inuoka"}
	fmt.Println("members are :", nekoma)
}
