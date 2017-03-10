package main

import(
	"os"
	"strconv"
	"bufio"
)

func main() {
	//set up user input stuff
	scanner := bufio.NewScanner(os.Stdin)

	//ask user for input
	print("Enter text: ")
	scanner.Scan()
	input := scanner.Text()

	//convert input to string
	//[TO-DO]: input verification (i.e. make sure they put in a number)
	number, _ := strconv.Atoi(input)

	//start the keithing
	check_keith(number, parse_num_to_list(number))
}

func parse_num_to_list(number int) []int{
	temp := strconv.Itoa(number)
	var output = make([]int, len(temp))
	for i, r := range temp{
		output[i], _ = strconv.Atoi(string(r))
	}
	return output
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
