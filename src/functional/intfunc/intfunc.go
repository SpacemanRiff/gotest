package intfunc

import(
	"strconv"
	"math"
	"fmt"
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
	for i := 3; i < int(math.Sqrt(float64(number))); i+=2{
		for(number % i == 0){
			output = append(output, i)
			number = number / i
		}
	}
	if(number > 2){
		output = append(output, number)
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

func Print_sums(list []int, sum int){
	for i, n := range list{
		if i < len(list)-1 {
			fmt.Print(strconv.Itoa(n) + " + ")
		}else{
			fmt.Print(strconv.Itoa(n) + " = ")
		}
	}
	fmt.Print(strconv.Itoa(sum) + "\n")
}

