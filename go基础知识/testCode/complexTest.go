package main

import "fmt"

/*
	The complex built-in function constructs a complex value from two
	floating-point values. The real and imaginary parts must be of the same
	size, either float32 or float64 (or assignable to them), and the return
	value will be the corresponding complex type (complex64 for float32,
	complex128 for float64)

	complex是复数， complex64和complex128，分别由float32和float64组成
*/
func main() {
	c := complex(1, 2)
	fmt.Println(c) //(1+2i)

	b := complex(2, 4)
	fmt.Println(b + c) //(3+6i)
	fmt.Println(b * c) //(-6+8i)
	fmt.Println(b / c) //(2+0i)
}
