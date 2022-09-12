package main

import "fmt"

type Animal struct {
	name  string
	age   int
	color string
}

func test() {
	animal1 := Animal{"cat", 20, "green"}
	animal2 := Animal{
		name:  "dog",
		age:   1,
		color: "blue",
	}
	fmt.Println(animal1)
	fmt.Println(animal2)
}
func main() {
	test()
}
