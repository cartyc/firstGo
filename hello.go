package main

import "./greeting"

func main() {

	//var s = greeting.Salutation{"Joe", "hello"}
	var s []int
	s = make([]int, 3)
	s[0] = 1
	s[1] = 10
	s[2] = 500

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Joe", "Hi"},
		{"Mary", "What is up?"},
	}

	slice = slice[0:2]

	greeting.Greet(slice, greeting.CreatePrintFunction("???"), true, 1)

}