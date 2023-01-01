package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	lower = 97
	upper = 65
)

func checkBp(pocket1 string, pocket2 string) int {
	var commonItem rune
	for _, item := range pocket1 {
		if checkInList(item, pocket2) {
			commonItem = item
		}
	}

	return getPriority(commonItem)
}

func checkInList(target rune, bp string) bool {
	re := regexp.MustCompile(string(target))
	return re.Match([]byte(bp))
}

func getPriority(item rune) int {
	if item >= upper && item < lower {
		return (int(item) - upper) + 27
	}

	return (int(item) - lower) + 1
}

func checkGroup(group []string) int {
	var commonItem rune
	first := group[0]
	for _, item := range first {
		if checkInList(item, group[1]) && checkInList(item, group[2]) {
			commonItem = rune(item)
		}
	}

	return getPriority(commonItem)
}

func main() {
	file, err := os.ReadFile("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bpList := strings.Split(string(file), "\n")

	var result int
	for _, bp := range bpList {
		if bp != "" {
			val := len(bp) / 2
			x := bp[:val]
			y := bp[val:]
			result += checkBp(x, y)
		}
	}
	fmt.Println("Part 1: ", result)

	var result2 int
	for i := 0; i < len(bpList); i += 3 {
		// Assume that list is always divisible by 3
		if i+3 < len(bpList) {
			group := bpList[i : i+3]
			result2 += checkGroup(group)
		}
	}

	fmt.Println("Part 2: ", result2)
}
