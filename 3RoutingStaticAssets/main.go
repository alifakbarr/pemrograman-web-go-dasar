package main

import "fmt"
import "net/http"

func main()  {
	http.Handle("/static/",
	http.StripPrefix("/static/",
	http.FileServer(http.Dir("assets"))))

	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		fmt.Println(err.Error())
	}
}