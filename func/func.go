package main

import "fmt"

func test(x, y int) int {
	return x + y
}

func test2(args ...int) int {
	sum := 0
	for _, value := range args {
		fmt.Println(value)
		sum += value
	}
	return sum
}

func test3(args ...string) {
	for _, value := range args {
		fmt.Println(value)
	}
}

func Test4() {
	fmt.Println("函数可见性")
	fmt.Println("函数名首字母大写则对所有包可见")
	fmt.Println("函数名首字母小写则对当前包可见")
}
func main() {
	fmt.Println(test(1, 2))
	fmt.Println(test2(1, 2, 3, 4, 5, 6, 7))
	var s []string
	s = append(s, []string{"hello", "world"}...)
	test3(s...)
}
