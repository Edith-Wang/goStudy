package main

import (
	"fmt"
	"unsafe"
)

/*
	uintptr is an integer type that is large enough to hold the bit pattern of any pointer.
	uintptr是一个整数类型, 其大小足以容纳任何指针
	即使uintptr变量仍然有效，由uintptr变量表示的地址处的数据也可能被GC回收
	uintptr 常用于与 unsafe.Pointer 打配合，用于做指针运算

	unsafe.Pointer是一个指针类型。
	但是unsafe.Pointer值不能被取消引用
	如果unsafe.Pointer变量仍然有效，则由unsafe.Pointer变量表示的地址处的数据不会被GC回收。

*/
func main() {
	//可以通过uintptr进行运算，计算出偏移量
	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])
	fmt.Printf("a[1] ptr: %v, a[3] ptr: %v\n", p1, unsafe.Pointer(&a[3])) //a[1] ptr: 0xc0000103a8, a[3] ptr: 0xc0000103b8
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[0]))
	fmt.Printf("p3 ptr: %v\n", p3) // p3 ptr: 0xc0000103b8a
	*(*int)(p3) = 6
	fmt.Println("a =", a) // a = [0 1 2 6]

	//通过计算偏移量来修改结构体中的值
	type Person struct {
		name   string
		age    int
		gender bool
	}

	who := Person{"John", 30, true}
	pp := unsafe.Pointer(&who)
	pname := (*string)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.name)))
	page := (*int)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.age)))
	pgender := (*bool)(unsafe.Pointer(uintptr(pp) + unsafe.Offsetof(who.gender)))
	*pname = "Alice"
	*page = 28
	*pgender = false
	fmt.Println(who) // {Alice 28 false}
}
