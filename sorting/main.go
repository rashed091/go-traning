package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	fruites := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruites))
	fmt.Println(fruites)

}
