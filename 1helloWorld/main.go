package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "welcome"

	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "hellow world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	// buat server
	address := "localhost:8000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
