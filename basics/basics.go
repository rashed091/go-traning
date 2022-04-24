package basics

import (
	"fmt"
	"reflect"
)

func main()  {
	var x = 4
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(float32(x) * 5.5))

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
