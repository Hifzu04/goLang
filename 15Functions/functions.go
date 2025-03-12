package main

import "fmt"

func adder(firstval int, secondval int) int {
	return firstval + secondval
}

// if i want want to add many values (uncounted)
func proadder(mydata ...int) int {
	total := 0
	for _, val := range mydata {
		total += val
	}
	return total
}

func main() {
	var a int = 10
	var b int = 12

	c := adder(a, b)
	fmt.Println(c) //22

	
	fmt.Println(proadder(1, 2, 3, 4, 5, 6))    //21

}
