package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	origin := "Hello, OTUS!"
	reversed := stringutil.Reverse(origin)
	fmt.Println(reversed)
}
