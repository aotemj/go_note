package main

import "fmt"

func test1(num int) {
	// go 中的 switch 不需要主动break
	switch num {
	case 1:
		fmt.Println("num =1")
	case 2:
		fmt.Println("num = 2")
	case 3:
		fmt.Println("num = 3")
	default:
		fmt.Println("num 不等于 1，2，3")
	}
}

func test2(num int) {
	switch num {
	case 1:
	case 2:
	case 3:
		fmt.Println("num = ", num)
	default:
		fmt.Println("num 不等于 1，2，3")
	}
}

func test3(num int) {
	switch num {
	case 1, 2, 3:
		fmt.Println("num = ", num)
	default:
		fmt.Println("num 不等于 1，2，3")
	}
}

func test4(num int) {
	// switch 也可以跟随判断语句
	switch {
	case num > 10:
		fmt.Println("num > 10")
	default:
		fmt.Println("num 不在条件范围内")
	}
}

func test5(s string) {
	//	在Go语言中 case 是一个独立的代码块，执行完毕后不会像C语言那样紧接着执行下一个 case，但是为了兼容一些移植代码，依然加入了 fallthrough 关键字来实现这一功能，代码如下
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}
}

func main() {
	//test1(3)
	//test1(4)
	//test2(3)
	//test2(4)
	//test3(3)
	//test3(4)
	//test4(3)
	//test4(4)
	//test4(41)

	test5("hello")
}
