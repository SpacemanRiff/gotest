package keith

import(
	. "./../intfunc"
	. "./../uinput"
	"fmt"
)

func main() {
	number := Get_int_put()

	Check_keith(number, true)
}

func Check_keith(check_val int, print_sums bool){
	check_keith(check_val, Parse_num_to_list(check_val),print_sums)
}

func check_keith(check_val int, num_list []int, print_sums bool){
	list_sum := 0
	for i := range num_list{
		list_sum = list_sum + num_list[i]
	}
	if print_sums{
		Print_sums(num_list, list_sum)
	}
	if(list_sum < check_val){
		check_keith(check_val, shift_list(num_list, list_sum), print_sums)
	} else if(list_sum == check_val){
		fmt.Printf("%v is a keith number\n", check_val)
	} else {
		if print_sums{
			fmt.Printf("%v is not a keith number\n", check_val)
		}
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
