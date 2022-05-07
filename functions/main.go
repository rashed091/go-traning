package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func vals() (int, int) {
	return 1, 5
}

func sums(nums ...int) {
	fmt.Print(nums, " : ")
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)

}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func fibb(n int) int {
	if n == 0 {
		return 1
	}
	return n * fibb(n-1)
}

func main() {
	res := plus(5, 8)
	fmt.Println(res)

	a, b := vals()
	fmt.Println(a, b)

	sums(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5}
	sums(nums...)

	nextint := intSeq()

	fmt.Println(nextint())

	fmt.Println(fibb(6))
}
