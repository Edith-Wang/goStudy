package main

import "fmt"

/*
	cap函数返回传入参数的容量, 不同的类型值不同：
	Array： 元素的数量，与len值相同
	Pointer to array： 元素的数量，与len值相同
	Slice： 重新切片时是slice最大长度，如果参数为nil时，容量为0
	Channel： 缓冲容量， 以元素为单位， 如果参数为nil时，容量为0
*/
func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("the Array size of capacity: %d\n", cap(arr))
	//输出结果: the Array size of capacity: 6

	slice := make([]int, 5, 8)
	fmt.Printf("the Slice size of capacity: %d\n", cap(slice))
	//输出结果: the Slice size of capacity: 8
	slice = nil
	fmt.Printf("the nil Slice size of capacity: %d\n", cap(slice))
	//输出结果: the nil Slice size of capacity: 0

	channel := make(chan int)
	fmt.Printf("the nil Channel size of capacity: %d\n", cap(channel))
	//输出结果: the Channel size of capacity: 0
	channel = make(chan int, 5)
	fmt.Printf("the Channel size of capacity: %d\n", cap(channel))
	//输出结果: the Channel size of capacity: 5
}
