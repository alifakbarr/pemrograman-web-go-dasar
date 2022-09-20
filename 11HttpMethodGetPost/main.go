package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("POST"))
		case "GET":
			w.Write([]byte("GET"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("server berjalan")
	http.ListenAndServe(":8080", nil)
}
