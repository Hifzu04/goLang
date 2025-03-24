// coverting slices of string into json
package main

import (
	"encoding/json"
	"fmt"
)



type students struct{
	Name string     // `json:"-"` //completely skip the name
	Dept string      `json:"departmentName"`        //in finalJson key will be departmetName instead of Dept
 	Spi float32
	OptedSubject []string  //slice of string
	isPassed bool
}




func encodeBigJson(){
	bigdata := []students{
		{"Afzal", "CSE" , 9.2 , []string{"DBMS","Computer Network"} , true},
		{"Hifzur" , "ECE" , 9 , []string{"DSA"} , false },
		{"Motsheem" , "EE" , 8.434 , nil , true},
}

    fmt.Printf("%+v\n" , bigdata)      //[{Name:Afzal Dapt:CSE Spi:9.2 OptedSubject:[DBMS Computer Network] isPassed:true} 
	                                 //  {Name:Hifzur Dapt:ECE Spi:9 OptedSubject:[DSA] isPassed:false} 
	                                  // {Name:Motsheem Dapt:EE Spi:8.434 OptedSubject:[] isPassed:true}]
  
	covnertedJson , _ := json.MarshalIndent(bigdata , "" , "\t")	
	fmt.Printf("%s\n" , covnertedJson)			//if fmt.Printf("%v\n" , convertedJson) , it will give databyte bcz Marshal convert the data in databye				  
    

  /*
  
 [
        {
                "Name": "Afzal",
                "departmentName": "CSE",
                "Spi": 9.2,
                "OptedSubject": [
                        "DBMS",
                        "Computer Network"
                ]
        },
        {
                "Name": "Hifzur",
                "departmentName": "ECE",
                "Spi": 9,
                "OptedSubject": [
                        "DSA"
                ]
        },
        {
                "Name": "Motsheem",
                "departmentName": "EE",
                "Spi": 8.434,
                "OptedSubject": null
        }
]
 
 
  */

}