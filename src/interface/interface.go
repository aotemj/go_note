package main

import "fmt"

// Animal 接口：
//  1. 接口只声明方法，不实现方法，结构体进行方法实现
//  2. 接口可以被多个结构体实现
type Animal interface {
	eat()
	sleep()
}

// 结构体1
type Cat struct {
	name string
}

func (c Cat) eat() {
	fmt.Println("cat is eating")
}

func (c Cat) sleep() {
	fmt.Println("cat is sleeping")
}

// 结构体2
type Dog struct {
	name string
}

func (d Dog) eat() {
	fmt.Println("Dog is eating")
}

func (d Dog) sleep() {
	fmt.Println("Dog is sleeping")
}

func main() {
	c := Cat{
		name: "cat1",
	}

	c.eat()
	c.sleep()

	d := Dog{
		name: "dog1",
	}
	d.eat()
	d.sleep()
}
