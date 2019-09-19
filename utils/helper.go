package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomArray(n, min, max int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		num := rand.Intn(max-min) + min
		arr = append(arr, num)
	}
	return arr
}
func TimeSpent(funcName string, inner func(arr []int), arr []int) {
	start := time.Now()
	inner(arr)
	fmt.Println(funcName+" time spent:", time.Since(start).Seconds())
}
