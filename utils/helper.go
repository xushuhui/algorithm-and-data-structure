package utils

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"reflect"
	"strings"
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

func GenerateOrderedArray(n int) []int{
	var arr []int
	for i := 0; i < n; i++ {

		arr = append(arr, i)
	}
	return arr
}
func GenerateTestOrderedArray(n int) []int {
	var arr []int
	for i := 0; i < n-1; i++ {

		arr = append(arr, i)
	}
	arr = append(arr, 1)
	return arr
}

// 判断arr数组是否有序
func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {

		if Compare(arr[i], arr[i+1]) > 0 {
			return false
		}
	}
	return true
}
func TimeSpent(funcName string, inner func(arr []int), arr []int) {
	start := time.Now()
	inner(arr)
	ts := time.Since(start).Seconds()
	n := len(arr)
	if !isSorted(arr) {
		return
	}
	fmt.Println(funcName, "n: ", n, " time :", ts)

}
func CopyArray(arr []int, n int) []int {
	var newArr = make([]int, n)
	copy(newArr, arr)
	return newArr
}

func Compare(a interface{}, b interface{}) int {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()

	if aType != bType {
		panic("cannot compare different type params")
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params")
	}
}
func ReadFile(filename string) []string {
	var words []string

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err != nil || io.EOF == err {
			break
		}

		wordSlice := strings.Fields(line)
		for _, word := range wordSlice {
			if word = extractStr(strings.ToLower(word)); word != "" {
				words = append(words, word)
			}
		}
	}

	return words
}
func extractStr(str string) string {
	var res []rune
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			res = append(res, letter)
		}
	}
	return string(res)
}
