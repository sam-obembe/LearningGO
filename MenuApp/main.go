package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espress", prices: map[string]float64{"single": 1.65, "double": 1.80, "triple": 1.95}},
	}

loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) Quit")

		in := bufio.NewReader((os.Stdin))

		choice, _ := in.ReadString('\n')

		choice = strings.TrimSpace(choice)
		switch choice {
		case "1":
			for _, val := range menu {
				fmt.Println(val.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, cost := range val.prices {
					fmt.Printf("\t%10s%10.2f\n", size, cost)
				}
			}
		case "2":
			fmt.Println("Please enter the name for the new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: name, prices: make(map[string]float64)})
		case "q":
			fmt.Println("terminating")
			break loop
			//terminate = true

		default:
			fmt.Println("No valid choice")
		}

	}

}
