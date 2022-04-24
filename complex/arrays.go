package complex

import "fmt"

func arrays()  {
	scores := [5]float64{9.2, 1.1, 3.3, 4.25, 3.7}
	fmt.Println(scores)
	fmt.Println(cap(scores))

	fruits := [5]string{"banana", "pear", "apple", "mango", "coconut"}
	selected := fruits[1:3]

	selected = append(selected, "jumbo")

	fmt.Println(selected)
	fmt.Println(len(selected))
	fmt.Println(cap(selected))
}
