package main

import "fmt"

/*
	len返回参数的长度, 根据类型返回值：
	Array: 返回的是数组的元素个数
	Pointer to array: 返回的是 *v 的元素个数，尽管v是nil也计算
	Slice, or map: 返回 v 的元素个数
	Channel: channel中队列元素（未读）的个数， 如果v是nil， 则返回0
*/
func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("arr length: %d\n", len(arr)) //arr length: 6

	slice := make([]int, 5, 8)
	fmt.Printf("slice length: %d\n", len(slice)) //slice length: 5

	maps := make(map[int]int, 5)
	maps[1] = 2
	maps[2] = 3
	fmt.Printf("map length: %d\n", len(maps)) //map length: 2

	channel := make(chan int, 10)
	fmt.Printf("channel length: %d\n", len(channel)) //channel length: 0
	channel <- 1
	channel <- 2
	fmt.Printf("channel length: %d\n", len(channel)) // channel length: 2
	<-channel
	fmt.Printf("channel length: %d\n", len(channel)) // channel length: 1
	close(channel)
}
