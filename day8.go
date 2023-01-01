package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertListToInt(str string) (result []int) {
	arr := strings.Split(str, "")
	for _, i := range arr {
		val, _ := strconv.Atoi(i)
		result = append(result, val)
	}
	return
}

func checkHorizontalVisibility(tree int, index int, row []int) bool {
	for i := 0; i < len(row); i++ {
		if row[i] >= tree && i < index {
			i = index
			continue
		}

		if i == index {
			return true
		}

		if row[i] >= tree && i > index {
			return false
		}

	}

	return true
}

func checkVerticalVisibility(tree int, treeIndex int, rowIndex int, rows [][]int) bool {
	for i := 0; i < len(rows); i++ {
		if rows[i][treeIndex] >= tree && i < rowIndex {
			i = rowIndex
			continue
		}

		if i == rowIndex {
			return true
		}

		if rows[i][treeIndex] >= tree && i > rowIndex {
			return false
		}

	}

	return true
}

func getVisibleTrees(row []int, rowIndex int, rows [][]int) int {
	result := 0

	for i, val := range row {
		if i == 0 || i == (len(row)-1) {
			result++
			continue
		}

		if checkHorizontalVisibility(val, i, row) {
			result++
			continue
		}

		if checkVerticalVisibility(val, i, rowIndex, rows) {
			result++
		}
	}

	return result
}

func getXScore(tree int, index int, row []int) (int, int) {
	left := 0
	right := 0

	for i := index - 1; i >= 0; i-- {
		left++
		if row[i] >= tree {
			break
		}
	}

	for i := index + 1; i < len(row); i++ {
		right++
		if row[i] >= tree {
			break
		}
	}

	return left, right
}

func getYScore(tree int, treeIndex int, rowIndex int, rows [][]int) (int, int) {
	top := 0
	bottom := 0

	for i := rowIndex - 1; i >= 0; i-- {
		top++

		if rows[i][treeIndex] >= tree {
			break
		}
	}

	for i := rowIndex + 1; i < len(rows); i++ {
		bottom++
		if rows[i][treeIndex] >= tree {
			break
		}
	}

	return top, bottom
}

func calculateScenicScore(row []int, rowIndex int, rows [][]int) int {
	result := 0
	for i, val := range row {
		if i == 0 || i == (len(row)-1) {
			continue
		}

		left, right := getXScore(val, i, row)
		top, bottom := getYScore(val, i, rowIndex, rows)
		sum := top * left * right * bottom
		if sum > result {
			result = sum
		}
	}

	return result
}

func main() {
	result := 0
	result2 := 0
	var rows [][]int
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := convertListToInt(scanner.Text())
		rows = append(rows, line)
	}

	for i, row := range rows {
		if i == 0 || i == (len(rows)-1) {
			result += len(row)
			continue
		}
		result += getVisibleTrees(row, i, rows)
		score := calculateScenicScore(row, i, rows)
		if score > result2 {
			result2 = score
		}
	}

	fmt.Println("Part 1:", result)
	fmt.Println("Part 2:", result2)
}
