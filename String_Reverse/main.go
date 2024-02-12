package main

import "fmt"

func main() {
	s := "aparnasukesh"
	fmt.Println(Reverse(s))
}

func Reverse(input string) string {
	rev := []byte(input)

	for i := 0; i < len(input)/2; i++ {
		rev[i], rev[len(rev)-i-1] = rev[len(rev)-i-1], rev[i]
	}
	return string(rev)
}
