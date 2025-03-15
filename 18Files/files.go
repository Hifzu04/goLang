package main

import (
	"fmt"
	"io"

	"os"
)

func main() {
	content := "this is my content that needs to in a  file , can i modify that"

	//creating a file
	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	//writting something into the created file
	//io.WriteString writes in  file a returns length of the content

	mylen, _ := io.WriteString(file, content)

	//return the length of content in the file
	fmt.Println(mylen)

	defer file.Close()

	//reading content from created file
	//io.Readfile() reads content in databyte instead of string format

	mydataByte, err := os.ReadFile("./myfile.txt")
	checkNilErr(err)
	// if err!= nil {
	// 	panic(err)
	// }

	fmt.Println(string(mydataByte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
