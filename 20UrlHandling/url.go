package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?courseName=react&paymentid=43blahblah"

func main() {

     //Parse parses a raw url into a [URL] structure.
	result , _ := url.Parse(myurl)
	
	fmt.Println(result.Scheme)      //https                                  
	fmt.Println(result.Host)         //lco.dev:3000  
	fmt.Println(result.Path)         //learn
	fmt.Println(result.RawQuery)       //courseName=react&paymentid=43blahblah
	fmt.Println(result.Port())         //3000


	// other way to store all the info of  the query of the url (note only query not host,scheme) in a better way


	qparams  := result.Query()

	// Query returns  in the form of : url.Values  ie key, val

	fmt.Printf("Query returns  in the form of : %T\n" ,qparams)

	fmt.Println(qparams["paymentid"])     //[43blahblah]
	fmt.Println(qparams["courseName"])     //[react]

	for _ , val := range qparams{
		fmt.Println("param is", val)          //param is [react]   param is [43blahblah]
	}




	



}
