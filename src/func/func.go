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

/*
多个参数类型不一致
当参数的类型不一致时，可以使用 interface 代替
*/
func test5(args ...interface{}) {
	for _, value := range args {
		switch value.(type) {
		case int:
			fmt.Println(value, "type is int")
			break
		case float32:
			fmt.Println(value, "type is float32")
			break
		case float64:
			fmt.Println(value, "type is float64")
			break
		case bool:
			fmt.Println(value, "type is bool")
			break
		case string:
			fmt.Println(value, "type is string")
			break
		case nil:
			fmt.Println(value, "type is nil")
			break
		default:
			fmt.Println(value, "type is unknown")
			break
		}
	}
}
func main() {
	fmt.Println(test(1, 2))
	fmt.Println(test2(1, 2, 3, 4, 5, 6, 7))
	var s []string
	s = append(s, []string{"hello", "world"}...)
	test3(s...)
	test5(1, "string", false)
}
