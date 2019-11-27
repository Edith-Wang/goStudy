package main

import "fmt"

/*
	append使用示例
	append属于内置函数，用于 slice 的元素添加操作
	如果添加元素时slice长度不够，可以自动为切片扩容
	通过以下示例可以看出，append函数可以将底层数组容量增长，一般增长为原来的2倍，扩容后会重新分配底层数组，再数据原来的复制
	有说1024以下是扩容两倍，但是在1024以上时，给的初始长度不同扩容的大小也不同，这个扩容的原理应该与slice底层实现有关，等学习到slice时再讨论
*/
func main() {

	slice := make([]int, 5, 8)
	fmt.Printf("slice ptr: %p, slice length: %d, slice capacity: %d\n", slice, len(slice), cap(slice))
	//输出结果： slice content: [0 0 0 0 0], slice ptr: 0xc00000e200, slice length: 5, slice capacity: 8
	fmt.Printf("slice[0] ptr: %p, slice[1] ptr: %p, slice[2] ptr: %p, slice[3] ptr: %p, slice[4] ptr: %p\n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4])
	//输出结果： slice[0] ptr: 0xc00000e200, slice[1] ptr: 0xc00000e208, slice[2] ptr: 0xc00000e210, slice[3] ptr: 0xc00000e218, slice[4] ptr: 0xc00000e220

	slice = append(slice, 1, 2)
	fmt.Printf("append slice ptr: %p,append slice length: %d,append slice capacity: %d\n", slice, len(slice), cap(slice))
	//输出结果： append slice content: [0 0 0 0 0 1 2],append slice ptr: 0xc00000e200,append slice length: 7,append slice capacity: 8
	fmt.Printf("slice[0] ptr: %p, slice[1] ptr: %p, slice[2] ptr: %p, slice[3] ptr: %p, slice[4] ptr: %p\n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4])
	//slice[0] ptr: 0xc00000e200, slice[1] ptr: 0xc00000e208, slice[2] ptr: 0xc00000e210, slice[3] ptr: 0xc00000e218, slice[4] ptr: 0xc00000e220

	slice = append(slice, 3)
	fmt.Printf("append2 slice ptr: %p,append2 slice length: %d,append2 slice capacity: %d\n", slice, len(slice), cap(slice))
	//输出结果： append2 slice content: [0 0 0 0 0 1 2 3],append2 slice ptr: 0xc00000e200,append2 slice length: 8,append2 slice capacity: 8
	fmt.Printf("slice[0] ptr: %p, slice[1] ptr: %p, slice[2] ptr: %p, slice[3] ptr: %p, slice[4] ptr: %p\n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4])
	//输出结果： slice[0] ptr: 0xc00000e200, slice[1] ptr: 0xc00000e208, slice[2] ptr: 0xc00000e210, slice[3] ptr: 0xc00000e218, slice[4] ptr: 0xc00000e220

	slice = append(slice, 4)
	fmt.Printf("append3 slice ptr: %p,append3 slice length: %d,append3 slice capacity: %d\n", slice, len(slice), cap(slice))
	//输出结果： append3 slice content: [0 0 0 0 0 1 2 3 4],append3 slice ptr: 0xc00007e100,append3 slice length: 9,append3 slice capacity: 16
	fmt.Printf("slice[0] ptr: %p, slice[1] ptr: %p, slice[2] ptr: %p, slice[3] ptr: %p, slice[4] ptr: %p\n", &slice[0], &slice[1], &slice[2], &slice[3], &slice[4])
	//输出结果： slice[0] ptr: 0xc00007e100, slice[1] ptr: 0xc00007e108, slice[2] ptr: 0xc00007e110, slice[3] ptr: 0xc00007e118, slice[4] ptr: 0xc00007e120

	testCapacityOfSlice()
}

func testCapacityOfSlice() {
	slice := make([]int, 1022, 1024)
	fmt.Printf("slice capacity: %d\n", cap(slice))
	//输出结果： slice capacity: 1024
	slice = append(slice, 1, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	fmt.Printf("slice append capacity: %d\n", cap(slice))
	//输出结果： slice append capacity: 1280

	slice1 := make([]int, 1023, 1024)
	fmt.Printf("slice1 capacity: %d\n", cap(slice1))
	//输出结果： slice capacity: 1024
	slice1 = append(slice1, 1, 2)
	slice1 = append(slice1, 3)
	slice1 = append(slice1, 4)
	fmt.Printf("slice1 append capacity: %d\n", cap(slice1))
	//输出结果： slice append capacity: 2048

	slice2 := make([]int, 2047, 2048)
	fmt.Printf("slice2 capacity: %d\n", cap(slice2))
	//输出结果： slice2 capacity: 2048
	slice2 = append(slice2, 1, 2)
	slice2 = append(slice2, 3)
	slice2 = append(slice2, 4)
	fmt.Printf("slice2 append capacity: %d\n", cap(slice2))
	//输出结果： slice2 append capacity: 2560
}
