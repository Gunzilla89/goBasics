package main

import (
	"fmt"

	"github.com/Gunzilla89/goBasics/myFunctions"
)

func main() {
	fmt.Println("Hello World")
	x := 2
	y := 3

	variable := myFunctions.Add(x, y)

	fmt.Println(variable)

}
