package main

import (
	"fmt"
	"sort"
)

func main() {

	//var fruitlist = []string{}  //declaration

	var fruitlist = []string{"apple", "tomato", "peach"}

	//fruitlist[idx] = "newfruit" // will not work

	//add values to the fruitlist
	fruitlist = append(fruitlist, "mango", "banana")
	fmt.Println(fruitlist) //[apple tomato peach mango banana]

	//	fruitlist= append(fruitlist[1:])
	fmt.Println(fruitlist) //[tomato peach mango banana]

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist) //[tomato peach]

	// other method
	highScores := make([]int, 4) //array type int of size 4
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	fmt.Println(highScores)   //[234 945 465 867]

	highScores  = append(highScores , 555, 666, 321)
	fmt.Println(highScores)       //[234 945 465 867 555 666 321]

	sort.Ints(highScores)
	fmt.Println(highScores)         // sorted ordr  [234 321 465 555 666 867 945]

}
