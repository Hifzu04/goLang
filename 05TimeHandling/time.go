package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println("curr time is ", presentTime)
	//curr time is  2025-02-16 17:57:45.744305305 +0530 IST m=+0.000026131
	// i want to format the time in my own way

	fmt.Println(presentTime.Format("01-02-2006"))  //02-16-2025  mm-dd-yy

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday")) //16-02-2025 18:02:37 Sunday  dd-mm-yy hh:mm:ss day

	//to create time 
	createdTime :=(time.Date(2049 , time.February , 15 , 12 , 59, 04, 0 , time.Local))

	fmt.Println(createdTime)  //2049-02-15 12:59:04 +0000 IST

	// i wamt tp format my created time  in this way month ,year

	fmt.Println(createdTime.Format("02 2006")) //15 2049
	



}
