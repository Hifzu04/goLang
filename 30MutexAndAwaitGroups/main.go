package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("if you are not using mutex you gonna get trapped in race condn");

	wg := &sync.WaitGroup{}
	mt := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3)

	go func(wg  *sync.WaitGroup , mt *sync.Mutex){
		
		fmt.Println("one r")
		mt.Lock()
		score = append(score, 1)
		mt.Unlock()
	}(wg ,mt)
	go func(wg *sync.WaitGroup, mt *sync.Mutex){
		
		fmt.Println("two r")
		mt.Lock()
		score = append(score, 2)
		mt.Unlock()
		wg.Done()
	}(wg,mt)
	go func(wg *sync.WaitGroup ,mt *sync.Mutex){
		
		fmt.Println("three r")
		mt.Lock()
		score = append(score, 3)
		mt.Unlock()
		wg.Done()
	}(wg ,mt)
	go func(wg *sync.WaitGroup,mt *sync.Mutex){
		
		fmt.Println("four r")
		mt.Lock()
		score = append(score, 4)
		mt.Unlock()
		wg.Done()
	}(wg ,mt)

	wg.Wait()
}