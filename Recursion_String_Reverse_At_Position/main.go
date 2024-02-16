package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aparna sukesh"
	fmt.Println(ReverseString(str))
}
func ReverseString(str string) string {
	words := strings.Split(str, " ")
	//var res string
	for i := 0; i < len(words); i++ {
		rev := Reverse(words[i], "", len(words[i])-1)
		// fmt.Println(rev)
		// res = res + rev + " "
		words[i] = rev

	}
	joinedString := strings.Join(words, " ")
	return joinedString // res

}

func Reverse(str, r string, i int) string {
	if i < 0 {
		return r
	}

	return Reverse(str, r+string(str[i]), i-1)
}
