package complex

import "fmt"

func printAge()	(ageOfRooh int, ageOfRafan int)  {
	ageOfRafan = 2
	ageOfRooh = 2
	return
}

func agg(values ...int)	{
	for _, value:= range values {
		fmt.Println(value)
	}
}


func complex() {
	age1, age2 := printAge()
	fmt.Println(age1, age2)
	fmt.Println(printAge())
	agg(1, 2, 3, 4)
}
