package main

import (
	"fmt"
	makeArr "mySrc"
	"time"
)

func main() {
	arr := makeArr.MakeArrs(10000000)
	start := time.Now()
	QuickSort(arr, 0, len(arr)-1)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func QuickSort(arr []int, left int, right int) {
	if right <= left {
		return
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
				Swap(arr, lt, gt)
			}
		}
		Swap(arr, left, lt)
		QuickSort(arr, left, lt-1)
		QuickSort(arr, lt+1, right)
	}
}

func Swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
