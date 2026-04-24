package main

import "fmt"

func greet(name string, age int) {
	fmt.Println("Hello,", name)
}

func main() {
	var name string
	var age int
	fmt.Println("Enter your name?")
	fmt.Scanln(&name)

	fmt.Println("Enter your age?")
	fmt.Scanln(&age)

	greet(name, age)
}
