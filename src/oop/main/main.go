package main

import stack "./../stack"
import singly "./../linkedlist/singly"
import "fmt"

func main() {
	stack_test()
	singly_linked_list_test()
}

func stack_test(){
	fmt.Println("started stack demonstation")
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
	fmt.Println()
}

func singly_linked_list_test() {
	fmt.Println("started singly linked list demonstration...")
	fmt.Println("adding nodes...")
	l := singly.New_List(1)
	fmt.Printf("first node: %v\n", singly.Get_Val_At(l, 0))
	for i := 0; i < 10; i++ {
		l = singly.Add_Node(l, 11 - i)
		fmt.Printf("%v\n", 11-i)
	}

	for i:= 0; i <10; i++ {
		fmt.Printf("Value at %v: %v\n", i, singly.Get_Val_At(l, i))
	}
}
