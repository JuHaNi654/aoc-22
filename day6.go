package main

import (
	"bytes"
	"fmt"
	"os"
)

func checkDublicate(signal string) bool {
	for _, c := range signal {
		size := bytes.Count([]byte(signal), []byte(string(c)))
		if size >= 2 {
			return true
		}
	}
	return false
}

func checkMarker(input []byte, size int) int {
	for i := 0; i < len(input); i++ {
		max := i + size
		if max < len(input) {
			signal := string(input[i:max])
			if !checkDublicate(signal) {
				return max
			}
		}
	}

	return 0
}

func main() {
	file, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println(err)
	}

	packetMarker := checkMarker(file, 4)
	messageMarker := checkMarker(file, 14)

	fmt.Println("Part 1: ", packetMarker)
	fmt.Println("Part 2: ", messageMarker)
}
