package main

import "fmt"

const myfrienname = "Siraj"    //const value

func main() {
	var mystr string = ("laraib")
	fmt.Println(mystr)

	//	fmt.Println("mystr is of type %T ", mystr)
	//mystr is of type %T  laraib
	//fmt.Println("mystr is of type %T \n",mystr)  will show error

	fmt.Printf("mystr is of type: %T \n", mystr)
	//laraib
	//mystr is of type: string

	var mybool bool = false //cant use 0 or 1
	fmt.Println(mybool)
	fmt.Printf("my bool is of type : %T\n", mybool)

	var myfloat float32 = 1323.439423402349234
	fmt.Println(myfloat) //op upto 8 bit 1323.4395 32/4(int 4 bit) = 8 bit
	fmt.Printf("my float is of typr : %T\n", myfloat)

	var anothervar int
	fmt.Println(anothervar) //op 0 (not any garbage value)

	fmt.Printf("my another var is of typr : %T\n", anothervar)

	var numbrofUser = 89080
	fmt.Println(numbrofUser) //wokks fine bcoz lexer internally defines variable data type

	numberofFollower := 100        //not allowed outside of the main
	fmt.Println(numberofFollower)



	fmt.Println(myfrienname)      //Siraj


	const Loginid = "fj;dafadkf"   //Public const Loginid  (as Capital L menas for public)
	fmt.Println(Loginid)

   


}
