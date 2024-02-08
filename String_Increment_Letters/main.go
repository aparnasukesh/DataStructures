package main

import "fmt"

func main() {
	name := "aparna"
	var key byte
	fmt.Printf("Enter the key")
	fmt.Scanf("%d", &key)
	newKey := key % 26
	fmt.Println(newKey)
	byt := []byte(name)
	for i := 0; i < len(byt); i++ {
		value := byt[i] + newKey
		if value <= 122 {
			byt[i] = value
		} else {
			byt[i] = 96 + (value % 122)
		}
	}
	fmt.Println(string(byt))
}
