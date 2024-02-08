package main

import (
	"fmt"
	"strings"
)

func main() {

	name := "Hello World"
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToUpper(string(name[1])))
	fmt.Println(name[0:1] + strings.ToUpper(string(name[1])) + name[1:])

}
