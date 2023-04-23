package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you like me to scream?")
	reader := bufio.NewReader(os.Stdin)

	screamVal, _ := reader.ReadString('\n')

	screamVal = strings.TrimSpace(screamVal)

	screamVal = strings.ToUpper(screamVal)

	fmt.Println(screamVal + "!")
}
