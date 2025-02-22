package main

import "fmt"

//how to remove a particular element from a slice

func main() {
	var courses = []string{"reactJs", "javaScript", "swift", "py", "ruby"}      //reactJs javaScript swift py ruby]
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)  //[reactJs javaScript py ruby]
   fmt.Println(courses)
}
