package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成长度为len的随机元素数组，元素范围为0-max
func genArray(len, max int) []int {
	arr := make([]int, len, len)
	rand.Seed(time.Now().Unix())
	for i, _ := range arr  {
		arr[i] = rand.Int() % max
	}
	return arr
}

//检验数组是否为单调递增
func checkSort(arr []int) {
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i] > arr[i+1] {
			fmt.Println("the result of sort is error.")
			return
		}
	}
	fmt.Println("the result of sort is right")
}
