package main

import "fmt"

func main() {
	x := 10
	y := 10

	//If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "green"
	if color == "red" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is NOT blue or red")
	}

	//Switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is red")
	default:
		fmt.Println("Color is NOT blue or red")
	}
}
