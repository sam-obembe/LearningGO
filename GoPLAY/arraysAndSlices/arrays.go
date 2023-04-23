package arraysandslices

import "fmt"

func DeclaringArrays() {
	var initArray [5]int

	fmt.Println(initArray)

	altArray := [3]int{1, 2, 3}
	fmt.Println(altArray)

	altArrayCopy := altArray
	altArray[0] = 0

	altArrayCopy[0] = -1

	fmt.Println(altArray)
	fmt.Println(altArrayCopy)
}
