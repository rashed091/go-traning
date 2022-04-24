package environment

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("Home: ", os.Getenv("HOME"))

	shell, ok := os.LookupEnv("SHELL")

	if !ok {
		fmt.Println("Not set")
	} else {
		fmt.Println("Shell: ", shell)
	}

}
