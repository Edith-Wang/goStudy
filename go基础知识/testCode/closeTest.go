package main

import "fmt"

/*
	The close built-in function closes a channel, which must be either
	bidirectional or send-only. It should be executed only by the sender,
	never the receiver, and has the effect of shutting down the channel after
	the last sent value is received. After the last value has been received
	from a closed channel c, any receive from c will succeed without
	blocking, returning the zero value for the channel element. The form
		x, ok := <-c
	will also set ok to false for a closed channel.

	close是用于关闭channel的内置函数, 这个channel必须是双向的或只写的， 这个方法只能由发送者调用，接受者不能，
	最后的发送值被接收后，关闭channel才管用。最后一个值从关闭的管道c中被接收后， 所有从c中接收值的接受者都不会发生阻塞，
	返回channel的零值， 如下代码：
		x, ok := <-c
	如果channl(c)已经关闭, ok返回false

	1. channel关闭后，还是可以从channel中取数据，在取完后，再取值时OK就返回false
	2. nil channel不能关闭
	3. channel不能重复关闭
	4. 只读的channel不能关闭

*/
func main() {
	channel := make(chan int, 5)
	for i := 0; i < 5; i++ {
		channel <- i
	}
	close(channel)
	for i := 0; i < 10; i++ {
		e, ok := <-channel
		fmt.Printf("channel index: %v, the result of ok: %v\n", e, ok)
		//channel index: 0, the result of ok: true
		//channel index: 1, the result of ok: true
		//channel index: 2, the result of ok: true
		//channel index: 3, the result of ok: true
		//channel index: 4, the result of ok: true
		//channel index: 0, the result of ok: false
		//channel index: 0, the result of ok: false
		//channel index: 0, the result of ok: false
		//channel index: 0, the result of ok: false
		//channel index: 0, the result of ok: false
		//if !ok {
		//	break
		//}
	}

	//nil channel不能关闭 执行报错:panic: close of nil channel
	//var nilChannel chan int
	//close(nilChannel)

	//重复关闭channel 执行报错: panic: close of closed channel
	//close(channel)

	//只读channel, 编译报错 invalid operation: close(channelReceive) (cannot close receive-only channel)
	//channelReceive := make(<-chan int, 1)
	//close(channelReceive)

	//关闭只写
	channelSend := make(chan<- int, 1)
	channelSend <- 1
	close(channelSend)

}
