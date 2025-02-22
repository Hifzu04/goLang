package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"
	languages["go"] = "golang"

	fmt.Println("list of all languages:", languages) //list of all languages: map[go:golang js:javascript py:python rb:ruby]

	fmt.Println(languages["py"]) //python

	delete(languages, "rb")
	fmt.Println(languages) //map[go:golang js:javascript py:python]

	//loops in a map
	for k, v := range languages {
		fmt.Println(k, v) //js javascript
		//py python
		//go golang
	}

	for key, val := range languages {
		fmt.Printf("for the key %v , value is %v\n", key, val)
		/* for the key js , value is javascript
		for the key py , value is python
		for the key go , value is golang   */
	}

	// i want to print all the value only , not key.
	for _, val := range languages {
		fmt.Printf("the value is %v\n", val)
		/*
	    the value is javascript
		the value is python
		the value is golang
		*/
	}
}
