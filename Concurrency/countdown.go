package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func countDown(count int) {
	for count >= 0 {
		fmt.Println(count)
		count--
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		countDown(5)
		wg.Done()
	}()

	go func() {
		countDown(6)
		wg.Done()
	}()

	wg.Wait()
}
