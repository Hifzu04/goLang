package main

import (
	"fmt"
	"io"

	"net/http"
)

func main() {
	url := "https://github.com/dashboard"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	mydataBytes, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(mydataBytes))

}
