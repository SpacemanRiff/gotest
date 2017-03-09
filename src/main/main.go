package main

import stack "./../stack"
import "fmt"

func main() {
	println("Started")
	s := stack.New_Stack(1)
	fmt.Println("Stack initialized")
	for i := 0; i < 10; i++ {
		s = stack.Push(s, i * (i + 1))
	}

	for stack.Has_Next(s) {
		peak_out := stack.Peak(s)
		var pop_out int
		pop_out, s = stack.Pop(s)
		fmt.Printf("peak: %v\n", peak_out)
		fmt.Printf("pop:  %v\n", pop_out)
	}
}
