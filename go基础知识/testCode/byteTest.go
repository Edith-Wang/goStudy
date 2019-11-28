package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

/*
	byte is an alias for uint8 and is equivalent to uint8 in all ways. It is
	used, by convention, to distinguish byte values from 8-bit unsigned
	integer values.
  	byte 其实就是uint8, 取值范围也和uint8一样, 0 ~ 255, 默认值为0, 用于区分字节值和8位无符号整数值
	这里提一下，golang中的byte与java中的byte取值范围不同
	java的byte定义为int8, 取值范围为-128 ~ 127
	如果要将byte在两种语言中转换，需要特殊处理， 参考：https://blog.csdn.net/hai046/article/details/52353963

	在Socket的Server和Client通信的过程中，传输的都是字节。而我们需要展示和使用的是字符串、整形等。这个时候，我们需要对字节进行处理，把byte类型的数据转成我们需要的类型。
	所以下面简单写几种转换的方式
*/
func main() {
	var a byte
	fmt.Println(a) //0

	//a = 256 //执行报错 constant 256 overflows byte

	bytesToInt32()
	bytesToInt16()
	int16ToBytes()
	byteToHex()
	hexToBye()
}

/*
	byte[] 与 int 互相转换
	字节与16位、32位及64位无符号整形之前的转换可使用“encoding/binary”包下的BigEndian（高位编址）与LittleEndian（低位编址）来操作，他们都实现了ByteOrder接口。
	提供了8位以上的无符号整形与byte数组之前的转换接口。
	type ByteOrder interface {
		Uint16([]byte) uint16
		Uint32([]byte) uint32
		Uint64([]byte) uint64
		PutUint16([]byte, uint16)
		PutUint32([]byte, uint32)
		PutUint64([]byte, uint64)
		String() string
	}

	即使用变量是int16,int32,int64定义的，也会调用Uint16, Uint32, Uint64实现ByteOrder的方法
	Read 中
		case *int16:
			*data = int16(order.Uint16(bs))
		case *uint16:
			*data = order.Uint16(bs)
		case *int32:
			*data = int32(order.Uint32(bs))
		case *uint32:
			*data = order.Uint32(bs)
		case *int64:
			*data = int64(order.Uint64(bs))
		case *uint64:
			*data = order.Uint64(bs)

	Write 中
		case *int16:
			order.PutUint16(bs, uint16(*v))
		case int16:
			order.PutUint16(bs, uint16(v))
		case *int32:
			order.PutUint32(bs, uint32(*v))
		case int32:
			order.PutUint32(bs, uint32(v))
		case *int64:
			order.PutUint64(bs, uint64(*v))
		case int64:
			order.PutUint64(bs, uint64(v))
	参考文献： https://studygolang.com/articles/6898
*/

func bytesToInt32() {
	a := []byte{12, 23, 34, 45, 2, 23, 1, 44, 55}
	buf := bytes.NewReader(a)
	var b int32
	//binary.BigEndian 最高位字节(按照从低地址到高地址的顺序存放数据的高位字节到低位字节)
	//binary.LittleEndian 最低位字节(按照从低地址到高地址的顺序存放据的低位字节到高位字节)
	binary.Read(buf, binary.BigEndian, &b)
	fmt.Println(b) //202842669
}

func bytesToInt16() {
	a := []byte{12, 23, 34, 45, 2, 23, 1, 44, 55}
	buf := bytes.NewReader(a)
	var b int16
	binary.Read(buf, binary.BigEndian, &b)
	fmt.Println(b) //3095
}

func int16ToBytes() {
	buf := bytes.NewBuffer([]byte{})
	b := int16(148)
	binary.Write(buf, binary.BigEndian, &b)
	fmt.Println(buf.Bytes()) //[0 148]
}

// []byte 与 16进制字符串
func byteToHex() {
	a := []byte{12, 23, 34, 45, 2, 23, 1, 44, 55}
	buffer := new(bytes.Buffer)
	for _, b := range a {

		s := strconv.FormatInt(int64(b&0xff), 16)
		if len(s) == 1 {
			buffer.WriteString("0")
		}
		buffer.WriteString(s)
	}
	fmt.Println(buffer.String())
}

//16进制字符串转[]byte
func hexToBye() {
	hex := "454449544857414e47"
	length := len(hex) / 2
	slice := make([]byte, length)
	rs := []rune(hex)

	for i := 0; i < length; i++ {
		s := string(rs[i*2 : i*2+2])
		value, _ := strconv.ParseInt(s, 16, 10)
		slice[i] = byte(value & 0xFF)
	}
	fmt.Println(slice)
}
