package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type warehouse struct {
	sections []*section
}

func newWarehouse() *warehouse {
	return &warehouse{}
}

func (wh *warehouse) scanCargoList(list map[int]string) {
	for section, cargo := range list {
		if cargo != "" {
			wh.registerCargo(cargo, section)
		}
	}
}

func (wh *warehouse) addSection(s *section) {
	wh.sections = append(wh.sections, s)
}

func (wh *warehouse) registerCargo(cargo string, id int) {
	for _, section := range wh.sections {
		if section.id == id {
			section.addCargo(cargo)
		}
	}
}

type section struct {
	id     int
	cargos []string
}

func newSection(id int) *section {
	return &section{
		id: id,
	}
}

func (s *section) addCargo(cargo string) {
	s.cargos = append(s.cargos, cargo)
}

func (s *section) appendCargo(cargos []string) {
	/*
		Uncomment code for part 1 solution
		var order []string
		for i := len(cargos); i > 0; i-- {
			order = append(order, cargos[i-1])
		}
	*/
	s.cargos = append(cargos, s.cargos...)
}

func (s *section) getCargo(amount int) []string {
	cargo := make([]string, len(s.cargos))
	copy(cargo, s.cargos)
	s.cargos = s.cargos[amount:]

	return cargo[:amount]
}

func generateCargoList(cargos []byte) map[int]string {
	cargoList := make(map[int]string)
	section := 1
	for i := 1; i < len(cargos); i += 4 {
		cargo := strings.ReplaceAll(string(cargos[i]), " ", "")
		cargoList[section] = cargo
		section++
	}

	return cargoList
}

func splitInput(file []byte) ([]byte, []byte) {
	input := bytes.Split(file, []byte("\n\n"))
	return input[0], input[1]
}

func operateCrane(cargo int, from int, to int, wh *warehouse) {
	var items []string

	for _, section := range wh.sections {
		if section.id == from {
			items = section.getCargo(cargo)
		}
	}

	for _, section := range wh.sections {
		if section.id == to {
			section.appendCargo(items)
		}
	}
}

func main() {
	file, err := os.ReadFile("./file.txt")
	if err != nil {
		fmt.Println(err)
	}

	whouse := newWarehouse()

	cargoBytes, craneBytes := splitInput(file)
	cargoRows := bytes.Split(cargoBytes, []byte("\n"))
	sections := cargoRows[len(cargoRows)-1:]
	cargoRows = cargoRows[:len(cargoRows)-1]

	x := strings.ReplaceAll(string(sections[0]), " ", "")

	// Generate sections
	for _, section := range x {
		id := string(section)
		val, _ := strconv.Atoi(id)

		if id != "" {
			sec := newSection(val)
			whouse.addSection(sec)
		}
	}

	// Assign cargos to the sections
	for _, row := range cargoRows {
		list := generateCargoList(row)
		whouse.scanCargoList(list)
	}

	// Operate crane
	input := bytes.Split(craneBytes, []byte("\n"))
	for _, cmd := range input {
		if string(cmd) != "" {
			split := strings.Split(string(cmd), " ")
			cargo, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])
			operateCrane(cargo, from, to, whouse)
		}
	}

	var result string
	// Print result
	for _, section := range whouse.sections {
		result += section.cargos[0]
	}

	fmt.Println("Result: ", result)
}
