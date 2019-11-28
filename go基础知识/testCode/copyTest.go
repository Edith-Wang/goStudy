package main

import "fmt"

/*
	The copy built-in function copies elements from a source slice into a
	destination slice. (As a special case, it also will copy bytes from a
	string to a slice of bytes.) The source and destination may overlap. Copy
	returns the number of elements copied, which will be the minimum of
	len(src) and len(dst).
	copy是将元素从源切片中复制到目标切片中。
	有一个特殊用途，可以将bytes从字符串复制到byte切片中。源和目标可能重叠。
	copy返回被复制元素的数量，取源切片长度和目标切片长度中的最小值

*/
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	length := copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Printf("slice1 is : %v, slice2 is : %v, the result of copy is %d\n", slice1, slice2, length)
	//输出结果:  slice1 is : [1 2 3 4 5], slice2 is : [1 2 3], the result of copy is 3

	length = copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Printf("slice1 is : %v, slice2 is : %v, the result of copy is %d\n", slice1, slice2, length)
	//输出结果: slice1 is : [5 4 3 4 5], slice2 is : [5 4 3], the result of copy is 3

	str := "aaaaaaaaa"
	b := make([]byte, 5)
	length = copy(b, str)
	fmt.Printf("b is : %v, the result of copy is %d\n", b, length)
	//输出结果: b is : [97 97 97 97 97], the result of copy is 5
}
