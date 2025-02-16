package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for giving an input,", input)

	fmt.Printf("the input is of type %T\n", input)
	// ConvertedRating, err := strconv.ParseFloat(input, 64)      //show err bcz str has a null char which it cant convert to float

	ConvertedRating, err := strconv.ParseFloat(strings.TrimSpace(input) , 64)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println("converted rating is", ConvertedRating)
		fmt.Printf("new converted string is of type %T\n" , ConvertedRating)

	}

}


// O/P
// 545
// thanks for giving an input, 545

// the input is of type string
// converted rating is 545
// new converted string is of type float64