package main

import (
	"fmt"
	"reflect"
)

// 结构体
type Animal struct {
	name  string
	age   int
	color string
}

func test() {
	// 声明结构体变量
	animal1 := Animal{"cat", 20, "green"}
	animal2 := Animal{
		name:  "dog",
		age:   1,
		color: "blue",
	}
	fmt.Println(animal1)
	fmt.Println(animal2)
}

func test2() {
	//	 结构体的比较
	animal1 := Animal{"cat", 20, "green"}
	animal2 := Animal{"cat", 20, "green"}
	fmt.Println(animal1 == animal2)                  // true
	fmt.Println(reflect.DeepEqual(animal1, animal2)) // true
}

// 结构体中添加方法
func (a Animal) test3() {
	fmt.Println("当前动物的名字是：", a.name)
	fmt.Println("当前动物的年龄是：", a.age)
}

// 注意： 如果要对结构体进行值的变更，必须传入当前结构体的指针
func (a *Animal) addAge(age int) {
	a.age += age
}

func main() {
	//test()
	//test2()
	animal1 := Animal{"cat", 20, "green"}
	animal1.test3()
	animal1.addAge(10)
	animal1.test3()
}
