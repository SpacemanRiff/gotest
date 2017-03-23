package main

import(
	"fmt"
)

func main(){
	test_strings := [4]string{"test", "hello, friend!", 
		"wow, what wonderful weather we're having", "ALL HAIL THE SYNTAX GODS"}
	for _,s := range test_strings{
		fmt.Println(s)
		fmt.Println(add_spaces(s))
	}
}

func add_spaces(s string) string{
	var app_str string
	for _,c := range s{
		app_str = app_str + string(c) + " "
	}
	return app_str
}
