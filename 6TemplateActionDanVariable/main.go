package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		person := Person{
			Name:    "Bruce Wayne",
			Gender:  "male",
			Hobbies: []string{"herex", "macol", "cangkrok"},
			Info:    Info{"Wayne Enterprice", "Gotham city"},
		}

		tmpl := template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server berjalan")
	http.ListenAndServe(":8080", nil)
}
