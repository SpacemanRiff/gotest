package main

import(
	"./../keith"
	//"./../smith"
)

func main(){
	for i := 0; i < 1000000; i++{
		keith.Check_keith(i, false)
		//smith.Check_smith(i, false)
	}
}
