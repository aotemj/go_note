package main

import "fmt"

// go 语言中的循环语句只有一个 for ,没有大多数语言中有的 while 、do while
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("当前循环中的索引值为：", i)
	}
}

func test1() {
	sum := 0
	for {
		sum++
		if sum > 100 {
			fmt.Println("当前值已不满足循环条件，主动break")
			break
		}
	}
	fmt.Println(sum)
}

func test2() {
	//	 go 中可以 主动break 已标记的循环
JLoop:
	for i := 0; i < 100; i++ {
		fmt.Println("i 的当前值为：", i)
		for j := 0; j < 100; j++ {
			fmt.Println("j 的当前值为：", j)
			if j > 50 {
				break JLoop
			}
		}
	}
}

func main() {
	//test()
	//test1()
	test2()
}
