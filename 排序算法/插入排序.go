package main

import (
	"fmt"
	makeArr "mySrc"
)

//1 9 2 4 3 5 7
//1     9 2 4 3 5 7
//1 9     2 4 3 5 7
//1 2 9    4 3 5 7
//1 2 4 9    3 5 7
//....
//1 2 3 4 5 7 9

//[1 9 9 4 3 5 7]
//[1 2 9 9 3 5 7]
//[1 2 4 9 9 5 7]
//[1 2 4 4 9 5 7]
//[1 2 3 4 9 9 7]
//[1 2 3 4 5 9 9]
//[1 2 3 4 5 7 9]
func main() {
	//arr := []int{1, 9, 2, 4, 3, 5, 7}
	arr := makeArr.MakeArrs(10)
	//start := time.Now()
	//insertSort(arr)
	//elapsed := time.Since(start)
	//fmt.Printf("执行时间： %v\n", elapsed)
	//10000 执行时间： 16.9526ms, 14.9584ms, 16.
	//insertSortBigToSmall(arr)
	insertSortRL(arr)
	fmt.Println(arr)
}

//从左到右 从小到大排序
func insertSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		v := arr[i]
		j := i - 1
		for j >= 0 && v < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
	}
	return arr
}

//从左到右 从大到小排序
func insertSortBigToSmall(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j := i - 1
		for j >= 0 && temp > arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = temp
	}
	return arr
}

//从右到左 从小到大排序
func insertSortRL(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := len(arr) - 2; i >= 0; i-- {
		temp := arr[i]
		j := i + 1
		for j <= len(arr)-1 && temp > arr[j] {
			arr[j-1] = arr[j]
			j++
		}
		arr[j-1] = temp
	}
	return arr
}

//从右到左 从大到小排序
func insertSortRLBigToSmall(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := len(arr) - 2; i > 0; i-- {
		temp := arr[i]
		j := i + 1
		for j <= len(arr)-1 && temp < arr[j] {
			arr[j-1] = arr[j]
			j++
		}
		arr[j-1] = temp
	}
	return arr
}
