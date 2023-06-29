package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// GetRequest()
	// PostJSONRequest()
	PostFormRequest()
}

func GetRequest() {
	const myURL = "http://localhost:8080/get"
	response, err := http.Get(myURL)
	checkNilErr(err)
	defer response.Body.Close()

	// // method 1
	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))

	// method 2
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	bytecount, _ := responseString.Write(content)
	fmt.Println(bytecount)
	fmt.Println(responseString.String())

}

func PostJSONRequest() {
	const myURL = "http://localhost:8080/post"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"course": "aao golang seekhein hum",
		"price": 0,
		"platform": "youtube"
	}
	`)

	response, err := http.Post(myURL, "application/json", requestBody)
	checkNilErr(err)
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PostFormRequest() {
	const myURL = "http://localhost:8080/postform"

	data := url.Values{}
	data.Add("firstname", "kakashi")
	data.Add("lastname", "hatake")
	data.Add("email", "kakashi@hokage.nin")

	response, err := http.PostForm(myURL, data)
	checkNilErr(err)
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
