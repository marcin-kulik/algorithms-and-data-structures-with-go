package main

import (
	"data-structures/queue"
	"fmt"
)

func main (){

	list := queue.NewFifoImpl()

	list.Push(1)
	list.Push(2)
	list.Push(3)
	fmt.Println("First pop ", list.Pop())

	list.Push(4)
	list.Push(2)
	fmt.Println("Second pop ", list.Pop())
	fmt.Println("3rd pop ", list.Pop())
	fmt.Println("4th pop ", list.Pop())

}