package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func readFileData(path string) []int {

	readFile, err := os.Open(path)
	if err != nil {
		fmt.Errorf("error reading %s: %v", path, err)
		os.Exit(1)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var numbers []int
	for fileScanner.Scan() {
		int, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			fmt.Errorf("error parsing integers: %v", err)
			os.Exit(1)
		}
		numbers = append(numbers, int)

	}
	return numbers
}
func part1(arr1, arr2 []int) {
	sort.Ints(arr1)
	sort.Ints(arr2)
	fmt.Println(arr2)

	var runningTotal float64
	for i := 0; i < len(arr1); i++ {
		runningTotal += math.Abs(float64((arr2[i] - arr1[i])))
	}
	fmt.Printf("%8f", runningTotal)
}

func main() {

	arr1 := readFileData("arr1")
	arr2 := readFileData("arr2")
	appearanceMap := make(map[int]int)
	for _, v := range arr2 {
		appearanceMap[v]++
	}

	similarityScore := 0
	for _, v := range arr1 {
		similarityScore += v * appearanceMap[v]
	}
	fmt.Println(similarityScore)

}
