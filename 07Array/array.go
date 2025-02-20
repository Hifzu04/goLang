package main

import "fmt"

func main()  {
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "tomato"
	fruitlist[3] = "peach"

	fmt.Println(fruitlist)    //[apple tomato  peach]  since fruitlist[2] is not present thats y there's a space b/w tomato and peach

	fmt.Println(len(fruitlist))    //4 

	var veglist = [3]string {"potato" , "beams" , "mashrooms" }
	fmt.Println(veglist)

}