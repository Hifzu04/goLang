package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to go module")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)

	//http.ListenAndServe(":4000" ,r)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("hey there mod useres")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from here </h1>"))
}
