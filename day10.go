package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	cycleStart = 20
	threshhold = 40
)

type Signal struct {
	x            int
	cycle        int
	combined     int
	strength     int
	sprite       *Sprite
	track        []string
	spriteTracks [][]string
}

type Sprite struct {
	idx  int
	size int
}

func newSprite() *Sprite {
	return &Sprite{
		idx:  0,
		size: 3,
	}
}

func (s *Signal) drawPixel(pos int) {
	s.track[s.sprite.idx+1] = "#"
}

func newTrack() []string {
	track := make([]string, 0, 20)
	for i := 0; i < 20; i++ {
		track = append(track, ".")
	}
	return track
}

func newSignal() *Signal {
	return &Signal{
		x:            1,
		cycle:        0,
		combined:     0,
		strength:     0,
		sprite:       newSprite(),
		track:        newTrack(),
		spriteTracks: [][]string{},
	}
}

func (s *Signal) register(cmd string, val int) {
	switch cmd {
	case "addx":
		for i := 0; i < 2; i++ {
			s.cycle++
			s.drawPixel(i)
			s.checkStrength(cmd, val)
		}
		s.x += val
		s.sprite.idx = val
	case "noop":
		s.cycle += 1
		s.checkStrength(cmd, val)
		s.drawPixel(1)
	default:
		return
	}
}

func (s *Signal) checkStrength(cmd string, val int) {
	if s.cycle >= cycleStart && s.cycle%threshhold == 20 {
		s.strength = (s.cycle * s.x)
		s.combined += s.strength

		s.spriteTracks = append(s.spriteTracks, s.track)
		s.track = newTrack()
		s.sprite.idx = 0
	}
}

func parseRow(row string) (string, int) {
	var value int
	split := strings.Split(row, " ")
	if len(split) > 1 {
		parsed, _ := strconv.Atoi(split[1])
		value = parsed
	}
	return split[0], value
}

func main() {
	signal := newSignal()
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newTrack())

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cmd, value := parseRow(scanner.Text())
		signal.register(cmd, value)
	}

	fmt.Println("Part 1:", signal.combined)
}
