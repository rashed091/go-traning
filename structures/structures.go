package structures

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Employee struct {
	ID     uint64
	Salary uint32
	Person User
}

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) getArea() int {
	return r.Length * r.Width
}

func main() {
	u := User{ID: 1, FirstName: "Md", LastName: "Rasheduzzaman", Email: "rashed091nsu@gmail.com"}

	fmt.Println(u.FirstName)

	var react Rectangle
	react.Length = 5
	react.Width = 4
	fmt.Println(react.Length * react.Width)

	var rr = new(Rectangle)
	rr.Length = 5
	rr.Width = 6
	fmt.Println(rr.getArea())

	emp := Employee{
		ID:     1234,
		Salary: 60000,
		Person: User{
			ID:        1,
			FirstName: "Md",
			LastName:  "Rasheduzzaman",
			Email:     "rashed091nsu@gmail.com",
		},
	}
	fmt.Println(emp)

	data, err := json.Marshal(u)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
	}

	fmt.Printf("JSON string: %s\n", data)

}
