package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

//form data : specially in case of image , variable, etc

func main() {
      formData()
}

func formData() {
	myData := url.Values{}
	myData.Add("firstName", "xyz")
	myData.Add("lastName", "pqr")

	res, err:= http.PostForm("https://jsonplaceholder.typicode.com/posts", myData)

	if (err!= nil){
		panic(err)
	}

	defer res.Body.Close()



	//check whether all the newly added data are posted in the sever or not
	content, _ := io.ReadAll(res.Body)

	fmt.Println(string(content))

}
