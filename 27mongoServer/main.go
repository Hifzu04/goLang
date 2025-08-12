package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hifzu04/mongoServer/27mongoServer/router"
)


func main(){
	fmt.Println("welcome to mongodb server")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":27017" , r))

	fmt.Println("listening and serving")
}