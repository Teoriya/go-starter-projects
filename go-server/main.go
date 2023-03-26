package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Invalid Method", 404)
		return
	}
	fmt.Fprintf(w, "Yolo Guys wassup , Sidhant Here")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", 404)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Invalid Method", 404)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "There was an error parsing your for %v", err)
		return
	}
	fmt.Fprintf(w, "Post submitted succesfully\n")
	name := r.FormValue("q")
	fmt.Fprintf(w, "Hi %v", name)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
