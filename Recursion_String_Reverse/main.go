package main

import "fmt"

func main() {
	s := "aparnasukesh"
	i := len(s) - 1
	fmt.Println(Reverse(s, "", i))
}

func Reverse(s, r string, i int) string {
	fmt.Println(r)
	if i < 0 {
		return r
	}
	return Reverse(s, r+string(s[i]), i-1)
}
