package main

import (
	"fmt"
	"io"
	"os"
)


func main(){
	content := "this is my content that needs to in a file"

	//creating a file 
	file  ,err := os.Create("./myfile.txt")

	if err!= nil {
		panic(err)
	}

	//writting something into the created file
	//io.WriteString writes in  file a returns length of the content
	
	mylen , _ :=io.WriteString(file , content)
	fmt.Println(mylen)

	defer file.Close()


	
}