package main

import "fmt"

// 使用 const 定义枚举类型
const (
	// 可以在 const() 中添加一个关键字 iota, 每行的iota都会累加1 第一行的默认值为 0
	// 如果是iota * 10 ,则后面每一个常量的值都会 乘以10
	CONST1 = iota * 10
	CONST2
	CONST3
)

func main() {
	fmt.Println(CONST1, CONST2, CONST3)
}
