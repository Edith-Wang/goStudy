package main

import (
	"fmt"
	makeArr "mySrc"
	"time"
)

//二分插入排序与快速排序结合使用
func main() {
	arr := makeArr.MakeArrs(10000)
	start := time.Now()
	SQuickSortCall(arr)
	elapse := time.Since(start)
	fmt.Printf("执行时间： %v\n", elapse)
}

func SQuickSortCall(arr []int) []int {
	if len(arr) < 100 {
		return SBinarySort(arr, 0, len(arr)-1)
	} else {
		SQuickSort(arr, 0, len(arr)-1)
		return arr
	}
}

func SBinarySort(arr []int, start int, end int) []int {
	if end-start <= 1 {
		return arr
	} else {
		for i := start + 1; i <= end; i++ {
			p := findSBinaryMidIndex(arr, start, i-1, i)
			if i != p {
				for j := i; j > p; j-- {
					SSwap(arr, j, j-1)
				}
			}
		}
		return arr
	}
}

func findSBinaryMidIndex(arr []int, start int, end int, curr int) int {
	if start >= end {
		if arr[start] > arr[curr] {
			return start
		} else {
			return curr
		}
	}
	mid := (start + end) / 2
	if arr[mid] > arr[curr] {
		return findSBinaryMidIndex(arr, start, mid, curr)
	} else {
		return findSBinaryMidIndex(arr, mid+1, end, curr)
	}
}

func SQuickSort(arr []int, left int, right int) {
	if right-left < 100 {
		SBinarySort(arr, left, right)
	} else {
		lt := left
		gt := right
		for lt < gt {
			for lt < gt && arr[gt] >= arr[left] {
				gt--
			}
			for lt < gt && arr[lt] <= arr[left] {
				lt++
			}
			if lt < gt {
				SSwap(arr, lt, gt)
			}
		}
		SSwap(arr, left, lt)
		SQuickSort(arr, left, lt-1)
		SQuickSort(arr, lt+1, right)
	}
}

func SSwap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
