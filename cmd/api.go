package main

import (
	"fmt"
	"log"
	"net/http"
	"octopus/pkg/api"
)

func main() {
	fmt.Println("API starting up...")
	r := api.Router()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
