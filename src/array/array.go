package main

import "fmt"

func test1() {
	// 声明时没有指定元素的值，则赋值为默认值，默认值为0（根据数据类型区别"0值"）
	var arr [5]int
	var strArr [5]string
	var floatArr [5]float32
	var boolArr [5]bool
	fmt.Println(arr)      //[0,0,0,0,0]
	fmt.Println(strArr)   // ["","","","",""]
	fmt.Println(floatArr) //[0,0,0,0,0]
	fmt.Println(boolArr)  // [false,false,false,false,false]
}

func test2() {
	// 声明时对数组进行初始化
	arr := [5]int{1, 2, 3, 4, 5}
	// 部分声明
	arr2 := [5]int{1, 2, 3}
	// 部分声明，针对单个 index
	arr3 := [5]int{1: 10, 4: 20}

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func test3() {
	// 不同长度的数组不管数据类型是否相同，他们的类型都不同，类型与数组长度相关
	arr := [5]int{1, 2, 3, 4, 5}
	arr2 := [10]int{}
	arr3 := [...]int{1: 10, 2: 30}
	fmt.Printf("arr 的类型是:%T\n", arr)
	fmt.Printf("arr2 的类型是:%T\n", arr2)
	fmt.Printf("arr3 的类型是:%T\n", arr3)
}
func test4() {
	//	 创建多维数组
	arr := [2][2]string{{"1", "2"}, {"3", "4"}}
	fmt.Println(arr)
}

func test5() {
	// 遍历列表
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for index, value := range arr {
	//	fmt.Printf("arr[%d]=%d\n", index, value)
	//}

	for _, value := range arr {
		fmt.Printf("arr: %d\n", value)
	}
}

func test6() {
	// go语音中的 数组是值类型，复制的值不会影响到原数组
	arr := [...]int{1, 2, 3, 5}
	copy := arr
	copy[0] = 2

	fmt.Println(arr)
	fmt.Println(copy)
}

func main() {
	//arrayByValue()
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	test6()
}
