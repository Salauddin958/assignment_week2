package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// tested in hackerrank environment
}

// https://www.hackerrank.com/challenges/array-left-rotation/problem
// solution
func rotateLeft(d int32, arr []int32) []int32 {
	if arr == nil || len(arr) == 0 {
		return arr
	}
	arr = reverse(arr, 0, len(arr)-1)
	arr = reverse(arr, 0, len(arr)-1-int(d))
	arr = reverse(arr, len(arr)-int(d), len(arr)-1)
	return arr
}

func reverse(arr []int32, start int, end int) []int32 {
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
	return arr
}

//https://www.hackerrank.com/challenges/2d-array/problem
// solution

func hourglassSum(arr [][]int32) int32 {
	if arr == nil || len(arr) == 0 || len(arr[0]) == 0 {
		return 0
	}
	var maxSum int32
	maxSum = -9999
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if i+2 < len(arr) && j+2 < len(arr[0]) {
				var sum int32
				sum = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] +
					arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return int32(maxSum)
}

//https://www.hackerrank.com/challenges/arrays-ds/problem
// solution
func reverseArray(arr []int32) []int32 {
	if arr == nil || len(arr) == 0 {
		return arr
	}
	start := 0
	end := len(arr) - 1
	for start < end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
	return arr

}

//https://www.hackerrank.com/challenges/find-the-median/problem
// solution
func findMedian(arr []int32) int32 {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	if len(arr)%2 != 0 {
		return arr[len(arr)/2]
	} else {
		return (arr[len(arr)/2] + arr[len(arr)/2-1]) / 2
	}
}

//https://www.hackerrank.com/contests/codart-2-0/challenges/word-count-1
// solution
func wordCount() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int(nTemp)

	var answer = make(map[string][]int)
	for i := 0; i < n; i++ {
		arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		for _, value := range arrTemp {
			if _, ok := answer[value]; !ok {
				answer[value] = []int{}
			}
			answer[value] = append(answer[value], i+1)
		}
	}
	keys := make([]string, 0, len(answer))
	for k := range answer {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Print(key)
		for _, v := range answer[key] {
			fmt.Print(" ")
			fmt.Print(v)
		}
		fmt.Println()
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
