//delays the execution of the part of code untill nearby fn/code are returned
//when multiple defer exist , works in LIFO ordr

package main

import "fmt"

func main() {

	defer fmt.Println("one")
	fmt.Println("two")

	defer fmt.Println("three")
	
	fmt.Println("hello")
	mydefer()

}

func mydefer(){
	defer fmt.Println()
	for i:=0 ; i<5 ; i++{
		defer fmt.Print(i)
		
	}
	

}

// two
// hello
// 43210
// three
// one
