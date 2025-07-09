package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is from Dockerized GoLang App!!!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on Port 8080")
	http.ListenAndServe(":8080", nil)
}
