package main

import "fmt"

func test() {
	arr := [5]int{1, 2, 3, 4, 5}
	// 从数组中取切片： 目标数组[切片开始索引:切片结束索引] 注意： 当前结束索引不包括在切片内容内
	s := arr[1:3] // [2,3]
	fmt.Println(arr)
	fmt.Println(s)
}

func test2() {
	//	使用 make 关键字创建切片
	s := make([]int, 3, 5) // 创建一个 []int 类型， 长度为3 ， 容量为5 的 切片
	fmt.Println(s)
}

func test3() {
	s := []int{}
	l := len(s) // 获取当前切片长度
	fmt.Println(s)
	fmt.Println(l)
}

func test4() {
	s := []int{}
	s = append(s, 1) // 追加元素
	fmt.Println(s)
}

func test5() {
	// 数组为值类型， 切片为引用类型，当修改切片的值时， 其宿主元素也会被修改
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:]
	s[0] = 3
	fmt.Println(arr) // [3,2,3,4,5]
	fmt.Println(s)   // [3,2,3,4,5]
}

func main() {
	//test()
	//test2()
	//test3()
	//test4()
	test5()
}
