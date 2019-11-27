package main

import (
	"fmt"
	"unsafe"
)

/*
	bool is the set of boolean values, true and false.
	bool就是布尔类型，值为true或false, 默认为false
*/
func main() {
	var b bool
	fmt.Println(b)                //false
	fmt.Println(unsafe.Sizeof(b)) // 字节为1

	//bool赋值方式
	b = true
	fmt.Println(b) // true

	//bool = 1 //编译错误 cannot use 1 (type int) as type bool in assignment

	//bool = 0 //编译错误 cannot use 0 (type int) as type bool in assignment

	b = (1 == 0)
	b = 1 == 0     //括不括号都行，括上看着明确一些
	fmt.Println(b) //

	b = (1 != 0)
	fmt.Println(b) // true

	//b = bool(1) //编译错误 cannot convert 1 (type untyped number) to type bool

	//查看源码会发现，true和false的定义为
	//true和false是两个非类型的布尔值
	//true  = 0 == 0 // Untyped bool.
	//false = 0 != 0 // Untyped bool.
	b = bool(false)  //false
	b = bool(0 != 0) //false
}
