package main

import stack "./../stack"
import "fmt"

func main() {
	fmt.Println("started")
	fmt.Println("pushing...")
	s := stack.New_Stack(1)
	for i := 0; i < 10; i++ {
		s = stack.Push(s, i * (i + 1))
		fmt.Printf("pushed: %v\n", stack.Peak(s))
	}

	fmt.Println("\npeaking & popping...")
	for stack.Has_Next(s) {
		peak_out := stack.Peak(s)
		var pop_out int
		pop_out, s = stack.Pop(s)
		fmt.Printf("peak: %v\n", peak_out)
		fmt.Printf("pop:  %v\n", pop_out)
	}
}
