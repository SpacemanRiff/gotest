package main

import(
	"strconv"
	. "./../intfunc"
	. "./../uinput"
)

func main() {
	number := Get_int_put()

	check_keith(number, Parse_num_to_list(number))
}

func check_keith(check_val int, num_list []int){
	list_sum := 0
	for i := range num_list{
		list_sum = list_sum + num_list[i]
	}
	output_sum_as_string(num_list, list_sum)
	if(list_sum < check_val){
		check_keith(check_val, shift_list(num_list, list_sum))
	} else if(list_sum == check_val){
		println(strconv.Itoa(check_val) + " is a keith number")
	} else {
		println(strconv.Itoa(check_val) + " is not a keith number")
	}
}

func shift_list(start_list []int, add_val int) []int{ //TO-DO
	var return_list = make([]int, len(start_list))
	for i := range start_list{
		if (i < len(start_list)-1){
			return_list[i] = start_list[i+1]
		} else {
			return_list[i] = add_val
		}
	}
	return return_list
}

func output_sum_as_string(num_list []int, sum_val int){
	for i := range num_list{
		if(i < len(num_list)-1){
			print(strconv.Itoa(num_list[i]) + " + ")
		} else {
			print(strconv.Itoa(num_list[i]) + " = ")
		}
	}
	println(sum_val)
}
