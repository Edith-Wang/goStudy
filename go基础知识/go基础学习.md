# GO语言基础

## 特色
* 简洁、快速、安全
* 并行、有趣、开源
* 内存管理、数组安全、编译迅速

## 基础知识

* 基础类型

| 数据类型 | 类型 | 字节大小 | 范围 |
|----|----|----|----|
| bool | 布尔型 | 1 | true/false |
| uint8 | 数值型 | 1 | 0 ~ 255 |
| uint16 | 数值型 | 2 | 0 ~ 65535 |
| uint32 | 数值型 | 4 | 0 ~ 4294967295 
| uint64 | 数值型 | 8 | 0 ~ 18446744073709551615 |
| uint | 数值型 | 4 或 8 | |
| int8 | 数值型 | 1 | -128 ~ 127 |
| int16 | 数值型 | 2 | -32768 ~ 32767 |
| int32 | 数值型 | 4 | -2147483648 ~ 2147483647 |
| int64 | 数值型 | 8 | -9223372036854775808 ~ 9223372036854775807 |
| int | 数值型 | 4 或 8 | |
| uintptr | 数值型 | 8 | 存放一个指针 |
| float32 | 浮点型 | 4 | IEEE-754 |
| float64 | 浮点型 | 8 | IEEE-754 |
| complex64 | 复数 | 4 | 实数和复数 |
| complex128 | 复数 | 8 | 实数和复数 |
| byte | 字节 | 1 | 常用于ASCII码字符，128个基本的字符，大小写字母，阿拉伯数字，英文标点符号，常用其他符号 |

* 关键字

| 关键字 | 示例 | 用途 |
|----|----| ---- |
| break | | 跳出当前循环 |
| default | | 一般与select或switch一起使用，默认执行子句 |
| func | | 函数声明 |
| interface | | 接口 |
| select | | 用来选择哪个case中的发送或接收操作可以被立即执行。它类似于switch语句，但是它的case涉及到channel有关的I/O操作 |
| case | | select或switch中的执行条件 |
| defer | | 函数执行结束时会被调用的 |
| go | | 并发执行 |
| map | | map类型 |
| struct | | 结构体 |
| chan | | channel管道 |
| else | | 一般与if连用 |
| goto | | 跳转语句 |
| package | | 包 |
| switch | | 选择执行那个操作 |
| const | | 常量声明 |
| fallthrough | | 如果case带有fallthrough，程序会继续执行下一条case,不会再判断下一条case的值 |
| if | | 判断条件 |
| range | | 从slice、map等结构中取元素 |
| type | | 定义类型 |
| continue | | 跳过本次循环 |
| for | | 循环 |
| import | | 引入包 |
| return | | 返回 |
| var | | 定义变量 |

* 预定义标识符

| 标识符 | 示例 | 使用 |
|----|----|----|
| append | [append示例](./testCode/appendTest.go) | 属于内置函数，用于 slice 的元素尾部添加操作 |
| bool, true, false | [bool示例](./testCode/boolTest.go) | 布尔类型 |
| byte | [byte示例](./testCode/byteTest.go) | 字节类型 |
| cap | [cap示例](./testCode/capTest.go) | 数组切片分配的空间大小 |
| close | [close示例](./testCode/closeTest.go) | 关闭channel |
| uint, uint8, uint16, uint32, uint64|  | 无符号整数 |
| int, int8, int16, int32, int64|  | 整数 |
| float32, float64 |  | 浮点数 |
| uintptr | [uintptr示例](./testCode/uintptrTest.go) | 整数类型, 其大小足以容纳任何指针 |
| complex, complex64, complex128 | [complex示例](./testCode/complexTest.go) | 复数类型 |
| real, imag | [real和imag示例](./testCode/realImagTest.go) | 返回复数的实数和虚数部分 |
| copy | [copy示例](./testCode/copyTest.go) | 将元素从源切片中复制到目标切片中 |
| iota | [iota示例](./testCode/iotaTest.go) | 常量计数器 |
| len | [len示例](./testCode/lenTest.go)| 返回参数的长度 |
| make | [make示例](./testCode/makeTest.go)| 分配和初始化slice, map 或 chan |
| new | [new示例](./testCode/newTest.go) | 分配空间 |
| nil | [nil示例](./testCode/nilTest.go) | 空值 |