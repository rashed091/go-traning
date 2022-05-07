package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var x = 4
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(float32(x) * 5.5))

	var a, b int = 1, 2
	fmt.Println(a, b)
	fmt.Println(math.Pow(float64(a), 5))

	const n = 3.1416
	fmt.Println(n)




	var name string = "Rashed"
	fmt.Println(name)

	if x > 4 {
		fmt.Println("x is bigger then 4")
	} else {
		fmt.Println("x is lesser or equal to 4")
	}

	switch name {
	case "Rashed":
		fmt.Println(name)
		// fallthrough
	case "Rooh":
		fmt.Println("Rooh")
	default:
		fmt.Println("End")
	}

}
