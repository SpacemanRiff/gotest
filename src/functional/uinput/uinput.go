package uinput

import(
	"os"
	"bufio"
	"strconv"
)

func Get_input() string{
	scanner := bufio.NewScanner(os.Stdin)
	print("Enter number: ")
	scanner.Scan()
	return scanner.Text()
}

func Get_int_put() int{
	output, _ := strconv.Atoi(Get_input())
	return output
}
