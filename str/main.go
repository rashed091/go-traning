package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	fmt.Println(strings.Contains("Tests", "es"))

	fmt.Println(strings.Count("Rashed", "a"))

	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	fmt.Println([]byte("Rashed"))

	fmt.Println(string([]byte{'t', 'a', 's'}))

	const s = "สวัสดี"

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
