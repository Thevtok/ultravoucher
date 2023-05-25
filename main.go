package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(words []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	groupedAnagrams := make([][]string, 0, len(anagramMap))
	for _, anagrams := range anagramMap {
		groupedAnagrams = append(groupedAnagrams, anagrams)
	}

	return groupedAnagrams
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	words := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	groupedAnagrams := groupAnagrams(words)

	fmt.Println("[")
	for i, group := range groupedAnagrams {
		fmt.Print("[")
		for j, word := range group {
			fmt.Printf("'%s'", word)
			if j < len(group)-1 {
				fmt.Print(",")
			}
		}
		fmt.Print("]")
		if i < len(groupedAnagrams)-1 {
			fmt.Println(",")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("]")
}
