package learnprimitives

import "fmt"

func SimpleDataTypes() {
	//Strings
	var name string = "Sam"

	var surname = "Obembe"

	middleName := "Eniola"

	fmt.Print(name + middleName + surname)

	//Numbers -> int, uint, float32, float64, complex64, complex128
	age := 27
	fmt.Println(age)

	//Booleans
	isHappy := true
	fmt.Println(isHappy)

	// Errors

	//type conversion
	income := 300

	dec := float32(income)

	fmt.Println(income)
	fmt.Println(dec)

	gdp := 3.5879 //float 64
	fmt.Println(gdp)

	e := int(gdp)
	fmt.Println(e)
}
