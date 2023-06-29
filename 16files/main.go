package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "RASENTARENSHURIKEN!"

	file, err := os.Create("./theFile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println(length)

	defer file.Close()

	ReadFile("./theFile.txt")
}

func ReadFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("content in databytes :", databyte)
	fmt.Println("content in string :", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
