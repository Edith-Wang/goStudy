package main

import "fmt"

/*
	make函数可以分配和初始化slice, map 或 chan
	和new相同之处是第一个参数传入type，不是value。 与new不同的是make返回类型，而不是指针
	根据类型返回不同的结果，如下：
	Slice: make(T, n, m)或make(T, n) T是类型,n长度,m容量, 如果m不传，容量和n相同，m不能小于n
	Map: make(T, n)或make(T) n是长度，空map分配了足够的空间来容纳指定数量的元素。如果不传n, 分配较小的起始大小
	Channel: make(T, n)或make(T) n是容量，channel初始化指定容量的缓冲区，如果是0或者未传n，无缓冲
*/
func main() {
	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 10)
	fmt.Println(slice1, len(slice1), cap(slice1)) //[0 0 0 0 0] 5 5
	fmt.Println(slice2, len(slice1), cap(slice2)) //[0 0 0 0 0] 5 10

	map1 := make(map[string]int)
	map2 := make(map[string]int, 5)
	fmt.Println(map1, len(map1)) //map[] 0
	fmt.Println(map2, len(map2)) //map[] 0

	channel1 := make(chan int)
	channel2 := make(chan int, 5)
	fmt.Println(channel1, cap(channel1)) //0xc000046060 0
	fmt.Println(channel2, cap(channel2)) //0xc000082000 5
}
