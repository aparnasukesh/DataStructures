package main

import "fmt"

func main() {
	str := "abcdfaghihjklmnkjhoijuii"
	// substrings abcdf
	// aghi
	// hjklmn
	fmt.Println(subString(str))
}

func subString(str string) string {
	res := make(map[string]bool)
	substr := ""
	sub := []string{}
	i := 0
	for i < len(str) {
		if _, ok := res[string(str[i])]; !ok {
			substr = substr + string(str[i])
			res[string(str[i])] = true
			i++
		} else {
			sub = append(sub, substr)
			substr = ""
			res = map[string]bool{}
		}
	}

	// Add the last substring to the slice
	if substr != "" {
		sub = append(sub, substr)
	}
	fmt.Println("sub: ", sub)
	// Find the largest substring
	maxlen := len(sub[0])
	ansStr := sub[0]
	for i := 1; i < len(sub); i++ {
		if len(sub[i]) > maxlen {
			ansStr = sub[i]
		}
	}

	fmt.Println(ansStr)

	// Concatenate the substrings into a single string with spaces
	result := ""
	for _, s := range sub {
		if result == "" {
			result = s
		} else {
			result = result + " " + s
		}
	}

	fmt.Println("result: ", result)
	return result
}
