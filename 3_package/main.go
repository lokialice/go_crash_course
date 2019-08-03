package main

import (
	"fmt"

	"github.com/lokialice/learnGO/go_crash_course/3_package/strutil"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Loki"))
	fmt.Println(getSum(3, 4))
	fmt.Println(strutil.Reverse("loki"))
}
