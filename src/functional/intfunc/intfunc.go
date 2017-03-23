package intfunc

import(
	"strconv"
)

func Parse_num_to_list(number int) []int{
	temp := strconv.Itoa(number)
	var output = make([]int, len(temp))
	for i, r := range temp{
		output[i], _ =  strconv.Atoi(string(r))
	}
	return output
}

func Fact_prime(number int) []int{
	var output = make([]int, 0)
	if(number%2 == 0){
		output = append(output, 2)
	}
	for i := 3; i < int(number/2); i+=2{
		for(number % i == 0){
			output = append(output, i)
			number = number / i
		}
	}
	return output
}

func Sum_list(list []int) int{
	output := 0
	for _, n := range list{
		output = output + n
	}
	return output
}

