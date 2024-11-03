package algsandds

// Implement an algorithm to see if a string has all unique characters
func IsUniqueString(text *string) bool {

	textDict := map[rune]int{}

	for _, chr := range *text {
		_, ok := textDict[chr]
		if ok {
			return false
		} else {
			textDict[chr] = 1
		}
	}

	return true
}
