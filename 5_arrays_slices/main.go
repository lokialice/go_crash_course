package main

import "fmt"

func main() {
	//Arrays
	var fruitsArr [2]string

	//Assign values
	fruitsArr[0] = "Apple"
	fruitsArr[1] = "Orange"

	fmt.Println(fruitsArr)
	fmt.Println(fruitsArr[0])

	//Declare and assign
	arr := [2]string{"Apple", "Orange"}
	fmt.Println(arr)

	arrSlices := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(len(arrSlices))
	fmt.Println(arrSlices[1:3])
}
