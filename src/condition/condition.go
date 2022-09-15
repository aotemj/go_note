package main

import "fmt"

func test1(num int) {
	if num > 10 {
		fmt.Println("num > 10")
	}
}

func test2(num int) {
	if num > 10 {
		fmt.Println("num >10")
	} else {
		fmt.Println("num<=10")
	}
}

func test3(num int) {
	if num > 10 {
		fmt.Println("num >10")
	} else if num > 5 {
		fmt.Println("5<num<=10")
	} else {
		fmt.Println("num<=5")
	}
}
func test4Connect(num int) (err string) {
	if num > 10 {
		err = "num 需小等于10"
	} else {
		err = ""
	}
	return err
}

// 特殊写法，缩小作用域 常用
func test4(num int) {
	if err := test4Connect(num); err == "" {
		fmt.Println("程序执行正常")
	} else {
		fmt.Println("程序异常 异常信息为：", err)
	}
}

func main() {
	//test1(10)
	//test2(10)
	//test3(10)
	//test3(1)
	test4(10)
	test4(15)
}
