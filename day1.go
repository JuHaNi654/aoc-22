package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sumGroup(arr string) (sum int) {
	for _, val := range strings.Split(arr, "\n") {
		x, _ := strconv.Atoi(val)
		sum += x
	}

	return
}

func getMaxValue(arr []int) (sum int) {
	for _, val := range arr {
		sum += val
	}
	return
}

func main() {
	file, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	groups := strings.Split(string(file), "\n\n")
	var values []int
	for _, s := range groups {
		values = append(values, sumGroup(s))
	}

	sort.Ints(values)

	fmt.Printf("Part1: %d\n\n", values[len(values)-1])

	fmt.Printf("Part2: %d\n\n", getMaxValue(values[len(values)-3:]))
}
