package dicts

import "fmt"

func DictDabble() {
	var saiyans = make(map[string]int)
	saiyans["Goku"] = 14000
	saiyans["Vegeta"] = 13000
	saiyans["Gohan"] = 11000
	saiyans["Trunks"] = 11000

	fmt.Println(saiyans)

	delete(saiyans, "Trunks")

	fmt.Println(saiyans)

	var jujutsu = map[string]string{
		"gojo":   "The Honored One",
		"Yuji":   "Semi-Grade 1",
		"Megumi": "Grade 2",
	}

	fmt.Println(jujutsu)

	var gojo, ispresent = jujutsu["gojo"]
	fmt.Printf("%s is present = %T", gojo, ispresent)
}
