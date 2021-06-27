package main

import (
	"fmt"
	"strings"
)

func main() {
	query := "de"
	strs := make([]string, 0)
	strs = append(strs, "dog", "deer", "deal", "deed", "asfg", "agsdopjiagish", "de de de de de ", "lllllllllllllllllllllllll")

	queryIndex := len(query)
	m := make(map[string][]string)

	for _, s := range strs {

		//just check in the slice ?
		if strings.Contains(s, query) {
			fmt.Print(s)
		}
		//make a map?
		m[s[:queryIndex]] = append(m[s[:queryIndex]], s)
	}
	fmt.Print(m)
}
