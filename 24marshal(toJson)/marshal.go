package main

import (
	"encoding/json"
	"fmt"
)


func main(){
    encodeJson()
	encodeBigJson()
}


//Task : to convert the following struct data into JSON

//DEFINING THE struct
type player struct {
	Name string
	Score int       
	Team string          `json:"team"`       //`json : "TeamName"`  // ❌ Incorrect syntax 
	IsPlaying bool
	//activeTeam []string     ////👉 Unexported(lowercase) eg(acitve team) fields are ignored by the json.Marshal() function.
	ActiveTeam []string
}


func encodeJson(){

	player1 := player{"Virat Kohli" , 65 , "india" , true , []string{"India", "RCB"}}

	finalJson,_ := json.Marshal(player1)

	fmt.Printf("%s\n",finalJson)        //[{"Name":"Virat Kohli","Score":65,"team":"india","IsPlaying":true,"ActiveTeam":["India","RCB"]}]
	fmt.Println(string(finalJson))
	
	fmt.Printf("%+v\n" , player1)     ////{Name:Virat Kohli Score:65 Team:india IsPlaying:true activeTeam:[India RCB]}

	IndentedJson,_ := json.MarshalIndent(player1,"", "\t")

	fmt.Println(string(IndentedJson))

	/*   indented Json
	[
        {
                "Name": "Virat Kohli",
                "Score": 65,
                "team": "india",
                "IsPlaying": true,
                "ActiveTeam": [
                        "India",
                        "RCB"
                ]
        }
]
	*/
	
}
