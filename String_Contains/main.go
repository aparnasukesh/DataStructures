package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Aparna Sukesh"
	sub := "Aparna"
	fmt.Println(strings.Contains(str, sub))
}
