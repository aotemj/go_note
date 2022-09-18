package main

import "fmt"

func test() func() {
	x := 10
	return func() {
		x = x + 10
		fmt.Println(x)
	}
}
func main() {
	a := test()
	a()
	a()
	a()
	b := test()
	b()
	b()
}
