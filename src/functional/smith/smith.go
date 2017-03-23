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
	var fact_dig_list = make([]int,0)

	for _, n := range fact_list{
		if n >= 10{
			temp_list := Parse_num_to_list(n)
			for _, m := range temp_list {
				fact_dig_list = append(fact_dig_list,m)
			}
		}else{
			fact_dig_list = append(fact_dig_list, n)
		}
	}

	num_sum := Sum_list(num_list)
	fact_sum := Sum_list(fact_dig_list)

	print_sums(num_list, num_sum)
	print_sums(fact_dig_list, fact_sum)
}

func print_sums(list []int, sum int){
	for i, n := range list{
		if i < len(list)-1{
			fmt.Print(strconv.Itoa(n) + " + ")
		}else{
			fmt.Print(strconv.Itoa(n) + " = ")
		}
	}
	fmt.Print(strconv.Itoa(sum) + "\n")
}
