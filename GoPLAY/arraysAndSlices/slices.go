package arraysandslices

import (
	"fmt"

	"golang.org/x/exp/slices"
)

// a slice is an abstraction that points to data in an array. It's a reference data type
func ManipulatingSlices() {
	var slice []int
	fmt.Println(slice) //nil

	slice = []int{1, 2, 3}
	fmt.Println("Slice after initialising", slice)
	fmt.Println("Original slice first item:", slice[0])

	slice[2] = 5

	slice = append(slice, 10, 11, 12)
	fmt.Println("Slice value after append: ", slice)

	copySlice := make([]int, len(slice), cap(slice)*2)

	copy(copySlice, slice)

	fmt.Println("Slice copy :", copySlice)

	slice = slices.Delete(slice, 2, 3)

	fmt.Println("Slice after delete:", slice)

	slice = append(slice, copySlice...)
	fmt.Println("Slice after second append", slice)
}
