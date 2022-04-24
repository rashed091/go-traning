package complex

import "fmt"

func maps() {
	users := make(map[int]string)

	users[0] = "rashed091nsu@gmail.com"
	users[1] = "roohesfar@gmail.com"

	fmt.Println(users)

	delete(users, 1)

	email, ok := users[1]
	fmt.Println(email, ok)
}
