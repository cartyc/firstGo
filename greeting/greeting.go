package greeting

import "fmt"

// Capital case indicates public availability lower case is private

type Salutation struct {
	Name string
	Greeting string
}

type Printer func(string) ()

func Greet(salutation []Salutation, do Printer, isFormal bool, times int){

	for _,s := range salutation{
		message, alternate := CreateMessage(s.Name, s.Greeting)

		for i := 0; i < times; i++{
			if prefix := getPrefix(s.Name); isFormal{
				do(prefix + message)
			} else {
				do(alternate)
			}
		}
	}


}

// name string =  arg; prefix string = named return
func getPrefix(name string) (prefix string){

	// Short hand declaration
	prefixMap := map[string]string{
		"Bob": "Mr ",
		"Joe": "Dr ",
		"Amy": "Dr ",
		"Mary": "Mrs ",
		}

	prefixMap["Joe"] = "jr "

	// delete key
	//delete(prefixMap, "Mary")

	// Long Hand declaration of map
	//var prefixMap map[string]string
	prefixMap = make(map[string]string)
	//
	//prefixMap["Bob"] = "Mr "
	//prefixMap["Joe"] = "Doc "
	//prefixMap["Amy"] = "Doc "
	//prefixMap["Mary"] = "Mrs "

	// Does it exist
	if value, exists := prefixMap[name]; exists {
		return value
	}

	return prefixMap[name]

	return
	}

func TypeSwitchTest(x interface{}) {
	 switch t := x.(type){
		case int: fmt.Println("int")
		case string: fmt.Println("string")
		case Salutation: fmt.Println("salutation")
		default: fmt.Printf("%T", t)
	}
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

