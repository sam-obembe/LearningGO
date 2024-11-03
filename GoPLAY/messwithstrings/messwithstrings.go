package messwithstrings

import "fmt"

//formatting io

func TryStringFormatting(word string) {

	//%b integer base 2
	//d integer base 10

	fmt.Printf("I am %v\n", 12)
	fmt.Printf("My name is %s\n", word)
	fmt.Printf("My bank account has $%f\n", 237.213)
	fmt.Printf("True or false, I have a great smile = %t\n", true)

	fmt.Printf("This value is of type %T\n", 223442)

	fmt.Printf("My exam score was 100%%\n")
}
