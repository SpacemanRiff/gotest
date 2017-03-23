package main

import "fmt"

func main(){
	for i := 0; true; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("london %v: ", i)
		}
		fmt.Print("\n")
	}
}
