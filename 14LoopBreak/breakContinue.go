package main

import "fmt"


func main(){
	days := []string{"sun" , "mon" , "tue" , "wed" , "thu" , "fri" , "sat"}
	fmt.Println(days)
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for i , dayname := range days{
		fmt.Printf("index is %v & day is %v\n" , i ,dayname)
		
	} 

	var myval = 1;
    
	//while loop
	for myval <10 {
		fmt.Printf("%d " , myval)        //1 2 3 4 5 6 7 8 9 
		myval++;
	}



 }