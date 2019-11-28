package main

import "fmt"

/*
	iota是常量计数器，只能在常量表达式中使用
	const出现时，iota被重置为0
	const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)

	参考文献： https://studygolang.com/articles/2192
*/

func main() {
	const a = iota                               // a=0
	fmt.Printf("default value of iota: %v\n", a) //default value of iota: 0
	const (
		b = iota //b=0
		c        //c=1
		d = 9
		e = 8
		f = iota
	)
	//输出结果： b is : 0, c is 1, d is 9, e is 8, f is 4
	fmt.Printf("b is : %v, c is %v, d is %v, e is %v, f is %v\n", b, c, d, e, f)
}

//定义枚举时比较方便
type Stereotype int

const (
	TypicalNoob           Stereotype = iota // 0
	TypicalHipster                          // 1
	TypicalUnixWizard                       // 2
	TypicalStartupFounder                   // 3
)

//可跳过的值
type AudioOutput int

const (
	OutMute   AudioOutput = iota // 0
	OutMono                      // 1
	OutStereo                    // 2
	_
	_
	OutSurround // 5
)

//位掩码表达式
type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

//定义量级
type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)
