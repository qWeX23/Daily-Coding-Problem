//Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters.
//For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	k := 5

	f, err := os.Open("..\\data\\words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	var submax, smax string
	var maxlen int
	for scanner.Scan() {
		s := scanner.Text()
		sub := doThings(s, k)
		if len(sub) > maxlen {
			submax = sub
			smax = s
			maxlen = len(sub)
		}

		i += 1
	}

	fmt.Print(smax)
	fmt.Print(" has max size substring of: ")
	fmt.Print(submax)

	fmt.Println()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func doThings(s string, k int) string {
	subStrings := make([]string, 0)

	for i := range s {

		for j := range s {
			if j+i > len(s) {
				break
			}
			sub := s[i : j+i]
			subStrings = append(subStrings, sub)

		}
	}

	mapDistincts := make(map[int]map[int]string)

	for _, x := range subStrings {
		i := MaxDistinctRunes(x)
		if mapDistincts[i] == nil {
			lenmap := make(map[int]string)
			lenmap[len(x)] = x
			mapDistincts[i] = lenmap
		} else {
			mapDistincts[i][len(x)] = x
		}

	}

	max := 0
	biggestString := ""
	for l, m := range mapDistincts {
		if l == k {
			for ln, str := range m {
				if ln > max {
					biggestString = str
				}
			}
		}
	}

	return biggestString
}

func MaxDistinctRunes(str string) int {
	m := make(map[rune]int)
	for _, x := range str {
		m[x] += 1
	}
	return len(m)
}
