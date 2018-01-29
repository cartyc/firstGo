package greeting

import "fmt"

// Capital case indicates public availability lower case is private

type Salutation struct {
	Name string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation Salutation, do Printer, isFormal bool){
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)
	if prefix := getPrefix(salutation.Name); isFormal{
		do(prefix + message)
	} else {
		do(alternate)
	}


}

// name string =  arg; prefix string = named return
func getPrefix(name string) (prefix string){
	switch name{
	case "Bob": prefix = "Mr "
	case "Joe": prefix = "Doctor "
	case "Mary": prefix = "Mrs "
	default: prefix = "Dude "
	}

	return prefix
}

func CreateMessage(name, greeting string) (message string, alternate string) {

	message = greeting + " " + name
	alternate = "Hey! " + name
	return
}

func Print(s string){
	fmt.Print(s)
}

func PrintLine(s string){
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string){
		fmt.Println(s + custom)
	}
}

func PrintCustom(s string, custom string){
	fmt.Println(s + custom)
}

func mathTest(number int) int {
	multiply := number * 2

	return multiply
}