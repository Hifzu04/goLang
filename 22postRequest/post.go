package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	//note server accepts JSON Data
	performPostRequest()

}

func performPostRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"

	//fake JSON payload
	requestbody := strings.NewReader(`
	{
	    "coursename" : "learnwithhrk",
		"price" : 0,
		"platform" : "xtz.in"

	}
	`)
	response ,err := http.Post(myurl,"application/json",requestbody )

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()

	//Read all the posted data
	content ,_ := io.ReadAll(response.Body)

	fmt.Println(string(content))
	//fmt.Println(response)



}
