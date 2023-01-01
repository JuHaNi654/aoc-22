package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toNum(arr []string) []int {
	result := make([]int, len(arr))
	for i, el := range arr {
		x, _ := strconv.Atoi(el)
		result[i] = x
	}

	return result
}

func getLongest(sec1 []int, sec2 []int) ([]int, []int) {
	range1 := (sec1[1] - sec1[0]) + 1
	range2 := (sec2[1] - sec2[0]) + 1
	if range1 > range2 {
		return sec1, sec2
	}

	return sec2, sec1
}

func checkContainingSections(sec1 []int, sec2 []int) bool {
	max, min := getLongest(sec1, sec2)
	return min[0] >= max[0] && min[1] <= max[1]
}

func checkOverlap(sec1 []int, sec2 []int) bool {
	max, min := getLongest(sec1, sec2)
	return (min[0] >= max[0] && min[0] <= max[1]) || (min[1] >= max[0] && min[1] <= max[1])
}

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	result := 0
	result2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sections := strings.Split(scanner.Text(), ",")
		n1 := toNum(strings.Split(sections[0], "-"))
		n2 := toNum(strings.Split(sections[1], "-"))
		if checkContainingSections(n1, n2) {
			result++
		}
		if checkOverlap(n1, n2) {
			result2++
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Part 1: ", result)
	fmt.Println("Part 2: ", result2)

}
