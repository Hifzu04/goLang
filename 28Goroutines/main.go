package main

import (
	"fmt"
	"time"
)

func speak(arg string) {

	for i := 0; i < 3; i++ {
		time.Sleep(3* time.Millisecond)
		fmt.Println(arg)
	}
}

func main() {

	go speak("hello world")
	
	speak("this is me")
	
}
