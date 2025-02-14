package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {

	fmt.Println("enter the rating of thr food")


    reader := bufio.NewReader(os.Stdin)
	

	

	input ,_ := reader.ReadString('\n')
	fmt.Println("thankss for rating:",input)
     fmt.Printf("type of rating is %T" , input)
}