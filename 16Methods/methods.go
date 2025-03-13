package main

import "fmt"

//when regular functions go into classes we call them methods
//used to accesss/manipulate the element from struct (object)

type User struct{
    Name string
    Email string
    Status bool
    Age int
    //these are objects
}

//Methods(to acesss the element)
func (u User) getStatus(){
    fmt.Println("is user active: " ,u.Status )
}
//method (to manipulate the element)
func (u User) newMail(){
      u.Email = "laraib@xyz.jmi"      //gives a copy of this mail. doesn't change the original email.
                                      //if i want to changae the original email, i have to make a pointer var
      fmt.Println(u.Email)
}





func main() {
     hrk := User{"hifzu" , "hifzu@xyz" , true , 18}
     fmt.Println(hrk)      //{hifzu hifzu@xyz true 18}
     fmt.Printf("%v\n" ,hrk) //{hifzu hifzu@xyz true 18}
     fmt.Printf("%+v\n" ,hrk) //{Name:hifzu Email:hifzu@xyz Status:true Age:18}

     fmt.Println(hrk.Email)  //hifzu@xyz



    hrk.getStatus()   //is user active:  true

    hrk.newMail()      //laraib@xyz.jmi





}
