package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
	byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
	used, by convention, to distinguish byte values from 8-bit unsigned
	integer values.
  	byte 其实就是uint8, 取值范围也和uint8一样, 0 ~ 255, 默认值为0, 用于区分字节值和8位无符号整数值
	这里提一下，golang中的byte与java中的byte取值范围不同
	java的byte定义为int8, 取值范围为-128 ~ 127
	如果要将byte在两种语言中转换，需要特殊处理， 参考：https://blog.csdn.net/hai046/article/details/52353963
*/
func main() {
	var a byte
	fmt.Println(a) //0

	//a = 256 //执行报错 constant 256 overflows byte

	bytesToInt32()
	bytesToInt16()
}

//byte数据 转 int
// bytes to int 32
func bytesToInt32() {
	a := []byte{12, 23, 34, 45, 2, 23, 1, 44, 55}
	buf := bytes.NewReader(a)
	var b int32
	//binary.BigEndian 最高位字节(按照从低地址到高地址的顺序存放数据的高位字节到低位字节)
	//binary.LittleEndian 最低位字节(按照从低地址到高地址的顺序存放据的低位字节到高位字节)
	binary.Read(buf, binary.BigEndian, &b)
	fmt.Println(b) //202842669
}

// bytes to int 16
func bytesToInt16() {
	a := []byte{12, 23, 34, 45, 2, 23, 1, 44, 55}
	buf := bytes.NewReader(a)
	var b int16
	binary.Read(buf, binary.BigEndian, &b)
	fmt.Println(b) //3095
}
