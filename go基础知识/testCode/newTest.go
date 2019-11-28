package main

import "fmt"

/*
	new函数用于分配内存。 第一个参数是type, 函数返回的是一个指向新分配的零值的指针
*/
func main() {
	s1 := new([]int)
	*s1 = append(*s1, 10)
	fmt.Println(s1, *s1) // &[10] [10]

	m1 := new(map[string]int)
	fmt.Println(m1, *m1) //&map[] map[]

	//用new新建channel没有容量，如果往channel send值，则报错fatal error: all goroutines are asleep - deadlock!
	x := new(chan int)
	//*x <- 1
	fmt.Println(x, *x) //0xc00008a028 <nil>
}
