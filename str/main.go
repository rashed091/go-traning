package main

import (
	"fmt"
	"strings"
)

func	main() {
	fmt.Println(strings.Contains("Tests", "es"))

	fmt.Println(strings.Count("Rashed", "a"))

	fmt.Println(strings.Join([]string{"a", "b"}, "-"))

	fmt.Println([]byte("Rashed"))

	fmt.Println(string([]byte{'t', 'a', 's'}))

}
