package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileData(path string) ([][]int, error) {

	readFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("error reading %s: %v", path, err)
		return nil, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines [][]int
	for fileScanner.Scan() {
		words := strings.Split(fileScanner.Text(), " ")
		var numbers []int
		for _, v := range words {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, n)
		}
		lines = append(lines, numbers)

	}
	return lines, nil
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func isAscending(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
		if (arr[i+1]-arr[i]) > 3 || (arr[i+1]-arr[i]) == 0 {
			return false
		}

	}
	return true
}
func isDescending(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
		if (arr[i+1]-arr[i]) < -3 || (arr[i+1]-arr[i]) == 0 {
			return false
		}
	}
	return true
}

func countTrues(arr []bool) int {
	sum := 0
	for _, v := range arr {
		if v {
			sum++
		}
	}
	return sum

}
func part2(lines [][]int) {

	valids := 0
	for _, v := range lines {
		if isAscending(v) || isDescending(v) {
			valids++
		} else {
			for j := 0; j < len(v); j++ {
				var newArr = make([]int, len(v))
				copy(newArr, v)
				copy(newArr[j:], newArr[j+1:])
				newArr = newArr[:len(newArr)-1]

				if isAscending(newArr) || isDescending(newArr) {
					valids++
					break
				}
			}
		}
	}
	fmt.Println(valids)
}

func main() {
	lines, err := readFileData("input.txt")
	if err != nil {
		fmt.Printf("%v", err)
	}
	part2(lines)
}
