package main

import (
	"fmt"
	"math"
	"unsafe"
)

// 有符号整型
func printInteger() {
	const num8 int8 = 127
	const num16 int16 = 32767
	const num32 int32 = math.MaxInt32
	const num64 int64 = math.MaxInt64
	var num int = math.MaxInt
	fmt.Printf("num8 的类型是 %T, num8 的大小是 %d, num8 是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16 的类型是 %T, num16 的大小是 %d, num16 是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32 的类型是 %T, num32 的大小是 %d, num32 是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64 的类型是 %T, num64 的大小是 %d, num64 是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num 的类型是 %T, num 的大小是 %d, num 是 %d\n", num, unsafe.Sizeof(num), num)
}

// 无符号整型
func unSingedInteger() {
	const num8 uint8 = 128
	const num16 uint16 = 32768
	const num32 uint32 = math.MaxUint32
	const num64 uint64 = math.MaxUint64
	const num uint = math.MaxUint

	fmt.Printf("num8 的类型是%T, num8 的大小是 %d, num8 是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16 的类型是%T, num16 的大小是 %d, num16 是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32 的类型是 %T, num32 的大小是 %d, num32 是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64 的类型是 %T, num64 的大小是 %d, num64 是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num 的类型是 %T, num 的大小是 %d, num 是 %d\n", num, unsafe.Sizeof(num), num)
}
func main() {
	fmt.Println("有符号整型：")
	printInteger()
	fmt.Println("============================================================")
	fmt.Println("无符号整型：")
	unSingedInteger()
}
