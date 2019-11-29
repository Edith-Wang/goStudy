package main

import (
	"math/rand"
	"time"
)

func MakeArr(length int) []int {
	var list []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(1000)))
	}
	return list
}
