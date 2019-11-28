package main

import "fmt"

/*
	The real built-in function returns the real part of the complex number c.
	The return value will be floating point type corresponding to the type of c.
	real 用于返回复数中实数部分

	The imag built-in function returns the imaginary part of the complex
	number c. The return value will be floating point type corresponding to
	the type of c.
	imag 用于返回复数中虚数部分

*/
func main() {
	b := complex(2, 4)

	fmt.Println(real(b)) //2
	fmt.Println(imag(b)) //4
}
