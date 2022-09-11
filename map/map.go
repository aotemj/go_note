package main

import "fmt"

func test() {
	//	使用make关键字声明map
	map1 := make(map[string]int)    // 其中 键类型为 string, 值类型为 int
	map2 := make(map[string]string) // 其中 键类型为 string, 值类型为 string
	fmt.Println(map1)
	fmt.Println(map2)
}

func test2() {
	//使用 := 声明 map
	map1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	map2 := map[int]int{
		1: 1,
		3: 4,
	}
	fmt.Println(map1)
	fmt.Println(map2)
}

func test3() {
	map1 := map[string]int{}
	map1["key1"] = 1
	// 检测当前元素是否存在于map中
	_, ok := map1["key2"]
	if !ok {
		map1["key2"] = 3
	}
	fmt.Println(map1)
	fmt.Println(len(map1))
	delete(map1, "key1")
	fmt.Println(map1)
	fmt.Println(len(map1))
}

func test4() {
	// map 为引用类型，当拷贝值被修改后，对应的原值也会被修改
	map1 := map[string]bool{
		"key1": true,
		"key2": false,
	}
	map2 := map1
	map2["key1"] = false
	fmt.Println(map1) // map[key1:false key2:false]
	fmt.Println(map2) // map[key1:false key2:false]
}

func test5() {
	map1 := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	fmt.Println(map1)

	// 遍历map
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}

func main() {
	//test()
	//test2()
	//test3()
	//test4()
	test5()
}
