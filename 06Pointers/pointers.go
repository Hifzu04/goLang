package main

import "fmt"


func main()  {
	var ptr *int 
	fmt.Println(ptr)   //<nil>

	myptr := 23 
	var addr = &myptr      //addr points to the addr of myptr
	fmt.Println(addr)    //0xc000098048    
	fmt.Println(*addr)    //23

	*addr = *addr+2
	fmt.Println(*addr)       //25 opertaion directly performed on myptr , not on the 'copy' of my number 
 
}