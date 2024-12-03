package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
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

func parseMul(group string, disable bool) (instruction, error) {
	var mult []int
	valid := true

	re := regexp.MustCompile(`\)`)
	group2 := re.ReplaceAllString(group, "")

	if disable {
		return instruction{
			opcode: "mul",
			loc:    []int{0, 0},
			args:   []int{0, 0},
		}, nil
	}
	//fmt.Printf("group: %s\n", group2)
	for _, num := range strings.Split(group2, ",") {
		operand, err := strconv.Atoi(num)
		if err != nil {
			valid = false
		}
		if valid {
			mult = append(mult, operand)
		}
	}
	if valid {
		return instruction{
			opcode: "mul",
			loc:    []int{0, 0},
			args:   mult,
		}, nil
	} else {
		return instruction{}, errors.New("malformed mult operator")
	}
}

type instruction struct {
	opcode string
	loc    []int
	args   []int
}

func main() {
	data := readFileData("input.txt")
	re := regexp.MustCompile(`((mul)|(don't)|(do))\(([(0-9),]*)\)`)

	matched := re.FindAllStringSubmatch(data, -1)
	var results []instruction
	disable := false
	for _, s := range matched {
		for groupIdx, group := range s {
			if group == "don't" {
				disable = true
			}
			if group == "do" {
				disable = false
			}

			if groupIdx == 0 {
				continue
			}
			ans, err := parseMul(group, disable)
			if err != nil {
				continue
			}
			results = append(results, ans)
		}
		fmt.Println(s)
	}

	sum := 0
	for _, v := range results {
		if v.opcode == "mul" {
			sum += v.args[0] * v.args[1]
		}
	}
	fmt.Printf("The sum of the mults is %d", sum)

}
