package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber) //generates a random number betweeen 1 to 6
	var mynum int

	switch diceNumber {
	case 1:
		mynum = 1
	case 2:
		mynum = 2
	case 3:
		mynum = 3
	case 4:
		mynum = 4
	case 5:
		mynum = 5
	case 6:
		mynum = 6
	default:
		fmt.Println("what was that")

	}
    

	if mynum == 1 {
		fmt.Printf("dice val is %v , you can  start or move %v\n" , mynum ,mynum);
	}else if mynum==6 {
		fmt.Printf("dice val is %v, you can move 6 move and get a new chance\n",mynum)
	}else{
		fmt.Printf("you can move %v move\n" , mynum)
	}


}
