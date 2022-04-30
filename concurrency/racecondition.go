package main

import (
	"fmt"
	"sync"
)

func main() {
	var memoryAccess sync.Mutex
	var data int

	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("If => The value is: ", data)
	} else {
		fmt.Println("Else => The value is: ", data)
	}
	memoryAccess.Unlock()
}
