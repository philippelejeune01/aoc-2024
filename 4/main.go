package main

import (
	"fmt"
	"os"
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

func xmasForward(x, y int, arr [][]rune) bool {
	if x >= len(arr)-3 {
		return false
	}
	return (arr[x][y] == 'X') && (arr[x+1][y] == 'M') && (arr[x+2][y] == 'A') && (arr[x+3][y] == 'S')
}

func xmasBackward(x, y int, arr [][]rune) bool {
	if x < 3 {
		return false
	}
	return (arr[x][y] == 'X') && (arr[x-1][y] == 'M') && (arr[x-2][y] == 'A') && (arr[x-3][y] == 'S')
}
func xmasDownward(x, y int, arr [][]rune) bool {
	if y >= len(arr)-3 {
		return false
	}
	return arr[x][y] == 'X' && arr[x][y+1] == 'M' && arr[x][y+2] == 'A' && arr[x][y+3] == 'S'
}
func xmasUpward(x, y int, arr [][]rune) bool {
	if y < 3 {
		return false
	}
	return arr[x][y] == 'X' && arr[x][y-1] == 'M' && arr[x][y-2] == 'A' && arr[x][y-3] == 'S'
}
func xmasUpandright(x, y int, arr [][]rune) bool {
	if (x >= len(arr)-3) || (y < 3) {
		return false
	}
	return arr[x][y] == 'X' && arr[x+1][y-1] == 'M' && arr[x+2][y-2] == 'A' && arr[x+3][y-3] == 'S'
}

func xmasDownandright(x, y int, arr [][]rune) bool {
	if (x >= len(arr)-3) || (y >= len(arr)-3) {
		return false
	}
	return arr[x][y] == 'X' && arr[x+1][y+1] == 'M' && arr[x+2][y+2] == 'A' && arr[x+3][y+3] == 'S'
}

func xmasUpandleft(x, y int, arr [][]rune) bool {
	if (x < 1) || (y < 1) || (y >= len(arr)-1) || (x >= len(arr)-1) {
		return false
	}
	return (arr[x][y] == 'A') && (arr[x-1][y-1] == 'M') && (arr[x+1][y+1] == 'S') || (arr[x][y] == 'A') && (arr[x-1][y-1] == 'S') && (arr[x+1][y+1] == 'M')
}
func xmasDownandleft(x, y int, arr [][]rune) bool {
	if x < 1 || (y >= len(arr)-1) || y < 1 || (x >= len(arr)-1) {
		return false
	}
	return (arr[x][y] == 'A' && arr[x-1][y+1] == 'M' && arr[x+1][y-1] == 'S') || (arr[x][y] == 'A' && arr[x-1][y+1] == 'S' && arr[x+1][y-1] == 'M')
}
func main() {

	data := readFileData("input2.txt")

	lines := strings.Split(data, "\n")

	//fmt.Println(lines)

	var charArr [][]rune
	for _, l := range lines {
		var line []rune
		for _, c := range l {
			line = append(line, c)

		}
		charArr = append(charArr, line)
	}
	charArr = charArr[:len(charArr)-1]

	numberOfXmas := 0
	for i, line := range charArr {
		if i == 0 {
			fmt.Printf("width of array : %d\n", len(charArr[i]))

		}
		for j := range line {

			// if xmasBackward(i, j, charArr) {
			// 	numberOfXmas++
			// }
			// if xmasForward(i, j, charArr) {

			// 	numberOfXmas++
			// }
			// if xmasUpward(i, j, charArr) {

			// 	numberOfXmas++
			// }
			// if xmasDownward(i, j, charArr) {

			// 	numberOfXmas++
			// }
			if xmasUpandleft(i, j, charArr) && xmasDownandleft(i, j, charArr) {

				numberOfXmas++
			}
			// if xmasUpandright(i, j, charArr) {

			// 	numberOfXmas++
			// }
			// if xmasDownandleft(i, j, charArr) {
			// 	numberOfXmas++
			// }
			// if xmasDownandright(i, j, charArr) {
			// 	numberOfXmas++
			// }

		}
	}

	fmt.Printf("height of array : %d\n", len(charArr))
	fmt.Printf("number of xmas: %d\n", numberOfXmas)

}
