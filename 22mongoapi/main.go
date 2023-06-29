package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hokagecv/mongoapi/router"
)

func main() {
	fmt.Println("server is getting started...")
	r := router.Router()
	fmt.Println("router created!")
	log.Fatal(http.ListenAndServe(":8000", r))
}
