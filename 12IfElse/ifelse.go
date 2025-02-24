package main

import (
	"fmt"
	
)

func main() {
	var num int
	fmt.Println("guess a random number")
	fmt.Scanf("%d", &num)
	var result string

	if num < 10 {
		result = "guess is too low"
	}else if num==15 {
		result = "you got it . wowwwwww"
	}else {
		result = "guess is too high"
	}

	fmt.Println(result)

}
