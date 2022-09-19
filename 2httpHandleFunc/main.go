package main

import "net/http"
import "fmt"


func main(){
	handlerIndex := func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("hello"))
	}

	handlerAbout := func (w http.ResponseWriter, r *http.Request)  {
		message := "Halaman About"
		w.Write([]byte(message))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/about", handlerAbout)

	fmt.Println("server started at localhost:8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil{
		fmt.Println(err.Error())
	}

}