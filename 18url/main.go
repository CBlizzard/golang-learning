package main

import (
	"fmt"
	"net/url"
)

const theUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghbl"

func main() {
	fmt.Println("full url is :", theUrl)

	// response, _ := url.Parse(theUrl)
	// fmt.Println("scheme is :", response.Scheme)
	// fmt.Println("host is :", response.Host)
	// fmt.Println("path is :", response.Path)
	// fmt.Println("query is :", response.RawQuery)

	// qparams := response.Query()
	// for _, param := range qparams {
	// 	fmt.Println("param is :", param)
	// }

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=kakashi",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("another url is :", anotherUrl)

}
