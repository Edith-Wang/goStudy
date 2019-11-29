package main

import (
	"fmt"
	makeArr "mySrc"
)

//用二分查找法找到插入的位置
func findMidIndex(arr []int, start int, end int, curr int) int {
	if start >= end {
		if arr[start] > arr[curr] {
			return start
		} else {
			return curr
		}
	}
	//二分查找递归
	mid := (start + end) / 2
	if arr[mid] > arr[curr] {
		return findMidIndex(arr, start, mid, curr)
	} else {
		return findMidIndex(arr, mid+1, end, curr)
	}
}

func main() {
	arr := makeArr.MakeArrs(10)
	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		p := findMidIndex(arr, 0, i-1, i)
		if i != p {
			temp := arr[i]
			for j := i; j > p; j-- {
				arr[j] = arr[j-1]
			}
			arr[p] = temp
		}
	}
	fmt.Println(arr)
}
