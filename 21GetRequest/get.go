package main

import (
	"fmt"
	"io"

	"net/http"
)
func main(){
	performgeGetRequest()
}




func performgeGetRequest()  {
	myurl := "https://go.dev/doc/effective_go"

	myResponse, _ := http.Get(myurl)
    
	//defer myResponse.Body.Close()


	fmt.Println(myResponse)
	fmt.Println("status code of my respone,",myResponse.StatusCode)
	fmt.Println("length of the content",myResponse.ContentLength)

	mycontent, _ := io.ReadAll(myResponse.Body)

	fmt.Println(string(mycontent))



	

}