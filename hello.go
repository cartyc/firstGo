package main

import "./greeting"

func main() {

	var s = greeting.Salutation{"Joe", "hello"}

	greeting.Greet(s, greeting.CreatePrintFunction("???"), true)

}