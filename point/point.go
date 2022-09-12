package main

import "fmt"

/**
关于指针的相关符号：
 &： 可以从一个变量中取得内存地址
 *： 赋值的左边，指该指针指向的变量，赋值的右边指 从一个指针变量中取得变量值（解引用）
*/

func test() {
	pStr := new(string)   // 创建指针并分配好地址
	*pStr = "hello world" // 给指针写入值
	fmt.Println(pStr)
	fmt.Println(*pStr)
}

func test1() {
	// 先声明一个指针变量，再从其他变量获取内存地址指针变量
	str := "string1"
	var p *string
	p = &str
	fmt.Println(p)
}
func test2() {
	str := "string1"
	pStr := &str // 获取当前 变量的指针地址
	fmt.Println(str)
	fmt.Println(pStr)
}

func test3() {
	str := "string2"
	pStr := &str
	fmt.Println("str = ", str)                // string2
	fmt.Println("*pStr = ", *pStr)            // string2
	fmt.Println("&str = ", &str)              // 0x....
	fmt.Println("pStr=", pStr)                // 0x....
	fmt.Println("pStr == &str", pStr == &str) // true
}

func test4() {
	myStr := "string"
	myInt := 1
	myBool := true
	myFloat := 0.11

	fmt.Printf("typeof myStr is %T\n", &myStr)
	fmt.Printf("typeof myInt is %T\n", &myInt)
	fmt.Printf("typeof myBool is %T\n", &myBool)
	fmt.Printf("typeof myFloat is %T\n", &myFloat)
}

func test5() {
	var pt *string // nil
	fmt.Println(pt)
	str := "string"
	pt = &str

	fmt.Println(pt) // 0xc000014250
}

func changeByPointer(value *int) {
	*value = 200
}

func test6() {
	myInt := 10
	pMyInt := &myInt
	// 执行 changeByPointer 之前
	fmt.Println("执行 changeByPointer 之前 pMyInt 为：", *pMyInt)
	changeByPointer(pMyInt)
	fmt.Println("执行 changeByPointer 之后 pMyInt 为：", *pMyInt)
}

func changeSlice(value []int) {
	value[0] = 10
}

func changeArray(value *[3]int) {
	(*value)[0] = 10
}

func test7() {
	// 指针的修改会影响到原始值
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[1:3]
	fmt.Println("改变之前的 s1 为：", s1)
	fmt.Println("改变之前的 arr 为：", arr)
	changeSlice(s1)
	fmt.Println("改变之后的 s1 为：", s1)
	fmt.Println("改变之后的 arr 为：", arr)
}

func main() {
	//test()
	//test3()
	//test4()
	//test5()
	//test6()
	test7()
}
