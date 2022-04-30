package concurrency

import "fmt"

func fib(max int, ch chan int)  {
	fi := make([]int, max)
	fi[0] = 0
	fi[1] = 1

	ch <- fi[0]
	ch <- fi[1]

	for i:= 2; i < max; i++ {
		fi[i] = fi[i-1] + fi[i-2]
		ch <- fi[i]
	}
	close(ch)
}


func main()  {
	ch := make(chan int)
	go fib(20, ch)

	for msg:= range ch {
		fmt.Println(msg)
	}
}
