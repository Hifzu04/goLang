package main

import "fmt"


type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("NO inheretence in go lang , no super or parent")
	hrk := user{"hifzur", "hifz.xyz@jmi", true, 19}     

	fmt.Println(hrk)      //{hifzur hifz.xyz@jmi true 19}
	

	fmt.Printf("%+v\n" , hrk)          //{Name:hifzur Email:hifz.xyz@jmi Status:true Age:19}

	fmt.Printf("name of the student is %v and email is %v\n" , hrk.Name , hrk.Email)   //name of the student is hifzur and email is hifz.xyz@jmi
	

}


