package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "abbcdddeee"
	//output:"eeedddbbca"

	fmt.Println(freqSort(str))
}

func freqSort(str string) string {
	freq := make(map[rune]int)

	for _, ch := range str {
		freq[ch]++
	}

	type pair struct {
		ch  rune
		fre int
	}

	pairs := make([]pair, 0, len(freq))

	for ch, f := range freq {
		pairs = append(pairs, pair{ch, f})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].fre > pairs[j].fre
	})
	res := "" // or res:=make([]rune,0,len(pairs))
	for _, p := range pairs {
		for i := 0; i < p.fre; i++ {
			res = res + string(p.ch) // or res=append(res,p.ch)
		}
	}

	return res // or string(res)
}
