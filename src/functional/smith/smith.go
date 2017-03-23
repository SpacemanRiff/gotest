package main

import(
	. "./../intfunc"
	. "./../uinput"
	"fmt"
	"strconv"
)

func main(){
	number := Get_int_put()

	check_smith(Parse_num_to_list(number), Fact_prime(number))
}

func check_smith(num_list []int, fact_list []int){
	num_sum := 0
	fact_sum := 0
	for i, n := range num_list{
		if i < len(num_list)-1{
			fmt.Print(strconv.Itoa(n) + " + ")
		}else{
			fmt.Print(strconv.Itoa(n) + " = ")
		}
		num_sum = num_sum + n
	}
	fmt.Print(strconv.Itoa(num_sum) + "\n")
	for i, n := range fact_list{
		if i < len(fact_list)-1{
			fmt.Print(strconv.Itoa(n) + " + ")
		}else{
			fmt.Print(strconv.Itoa(n) + " = ")
		}
		fact_sum = fact_sum + n
	}
	fmt.Print(strconv.Itoa(fact_sum) + "\n")
}
