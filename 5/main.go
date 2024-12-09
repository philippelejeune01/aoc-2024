package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFileData(path string) string {
	readFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading %s: %v\n", path, err)
		os.Exit(1)
	}

	return string(readFile)
}

type Rule struct {
	First  int
	Follow int
}

func parseRule(s string) Rule {
	parts := strings.Split(s, "|")
	k, _ := strconv.Atoi(parts[0])
	v, _ := strconv.Atoi(parts[1])
	return Rule{First: k, Follow: v}
}

func main() {
	data := readFileData("input.txt")

	parts := strings.Split(data, "\n\n")

	ruleStr := parts[0]
	updateStr := parts[1]
	var rules []Rule
	var updates [][]int
	for _, v := range strings.Split(ruleStr, "\n") {
		r := parseRule(v)
		rules = append(rules, r)
	}
	for _, k := range strings.Split(updateStr, "\n") {
		var nums []int
		for _, v := range strings.Split(k, ",") {
			n, _ := strconv.Atoi(v)
			nums = append(nums, n)
		}
		updates = append(updates, nums)
	}
	updates = updates[:len(updates)-1]
	fmt.Println(updates)
}
