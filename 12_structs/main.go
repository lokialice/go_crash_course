package main

import (
	"fmt"
	"strconv"
)

//Define Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender != "M" {
		p.lastName = spouseLastName
	} else {
		return
	}
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Sama", lastName: "Lich", city: "HCM", gender: "M", age: 24}

	//Alternative
	person2 := Person{"Loki", "Alice", "HN", "F", 32}

	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.lastName)
	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

}
