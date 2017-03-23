package main

import(
	. "./../intfunc"
	. "./../uinput"
	"fmt"
)

func main(){
	number := Get_int_put()

	Check_smith(number)
}

func Check_smith(number int){
	num_list := Parse_num_to_list(number)
	fact_list := Fact_prime(number)
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

	Print_sums(num_list, num_sum)
	Print_sums(fact_dig_list, fact_sum)
	if num_sum == fact_sum{
		fmt.Printf("%v is a smith number.\n", number)
	}else{
		fmt.Printf("%v is not a smith number.\n", number)
	}
}
