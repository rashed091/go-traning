package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {
	var x list.List
	x.PushBack(5)
	x.PushBack(2)
	x.PushBack(3)

	sort.Sort(x)

	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}
