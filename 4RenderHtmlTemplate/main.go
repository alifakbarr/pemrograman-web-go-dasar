package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func main() {
	// handle css
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := path.Join("views", "index.html")
		tmpl, err := template.ParseFiles(filepath)
		if err != nil {
			fmt.Println(err.Error())
		}

		data := map[string]interface{}{
			"title": "learning golang web",
			"name":  "batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("sudah jalan")
	http.ListenAndServe(":8080", nil)
}
