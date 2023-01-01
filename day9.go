package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type knot struct {
	x      int
	y      int
	coords map[[2]int]bool
	prev   *knot
	next   *knot
}

type bridge struct {
	head *knot
	tail *knot
}

func newBridge() *bridge {
	return &bridge{}
}

func newKnot(n *knot, p *knot) *knot {
	return &knot{
		x:      0,
		y:      0,
		coords: make(map[[2]int]bool),
		next:   n,
		prev:   p,
	}
}

func (b *bridge) createKnots(size int) {
	for i := 0; i < size; i++ {
		knot := newKnot(b.head, nil)

		if b.head != nil {
			b.head.prev = knot
		}

		b.head = knot

		lastKnot := b.head
		for lastKnot.next != nil {
			lastKnot = lastKnot.next
		}

		b.tail = lastKnot
	}
}

func isInRange(k1 *knot, k2 *knot) bool {
	return (((k1.x-k2.x <= 1) && (k1.x-k2.x >= -1)) &&
		((k1.y-k2.y <= 1) && (k1.y-k2.y >= -1)))
}

func (b *bridge) moveVertically(d int, s int) {
	for i := 0; i < s; i++ {
		knot := b.head
		knot.y += d
		for knot.next != nil {
			knot = knot.next
			if isInRange(knot, knot.prev) == false {
				knot.x = knot.prev.x + d
				knot.y = knot.prev.y
				if knot.next == nil {
					knot.coords[[2]int{knot.x, knot.y}] = true
				}
			}
		}
	}
}

func (b *bridge) moveHorizontally(d int, s int) {
	for i := 0; i < s; i++ {
		knot := b.head
		knot.x += d
		for knot.next != nil {
			knot = knot.next
			if isInRange(knot, knot.prev) == false {
				knot.x = knot.prev.x - d
				knot.y = knot.prev.y

				if knot.next == nil {
					knot.coords[[2]int{knot.x, knot.y}] = true
				}
			}
		}
	}
}

func (b *bridge) setCoords(d string, s int) {
	switch d {
	case "U":
		b.moveVertically(-1, s)
	case "D":
		b.moveVertically(1, s)
	case "L":
		b.moveHorizontally(-1, s)
	case "R":
		b.moveHorizontally(1, s)
	default:
		return
	}
}

type step struct {
	direction string
	step      int
}

func newStep(d string, s int) *step {
	return &step{
		direction: d,
		step:      s,
	}
}

func parseRow(row string) (string, int) {
	split := strings.Split(row, " ")
	step, _ := strconv.Atoi(split[1])
	return split[0], step
}

func main() {
	bridge := newBridge()
	bridge.createKnots(2)
	var stepList []*step
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		direction, step := parseRow(scanner.Text())

		val := newStep(direction, step)
		stepList = append(stepList, val)
	}

	for _, val := range stepList {
		bridge.setCoords(val.direction, val.step)
	}

	result := 0
	for _, visited := range bridge.tail.coords {
		if visited {
			result++
		}
	}
	fmt.Println("Part 1: ", result)

}
