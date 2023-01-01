package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	MAX_SIZE = 100_000
	UPDATE   = 30_000_000
	SYSTEM   = 70_000_000
)

type filesystem struct {
	root         *folder
	activeFolder *folder
}

func createNewFilesystem() *filesystem {
	return &filesystem{}
}

func (fs *filesystem) goToFolder(target string) {
	// Create root folder
	if fs.root == nil {
		folder := createNewFolder(target, nil)
		fs.root = folder
		fs.activeFolder = folder
		return
	}

	if target == ".." {
		folder := fs.activeFolder.parentDir
		fs.activeFolder = folder
		return
	}

	folder := fs.activeFolder.getFolder(target)
	fs.activeFolder = folder
}

type folder struct {
	parentDir *folder
	dir       string
	folders   []*folder
	files     []*file
	size      int
}

func (f *folder) getFolderSizeList() []int {
	sizeList := []int{}

	if (SYSTEM - f.size) > UPDATE {
		sizeList = append(sizeList, f.size)
	}

	for _, folder := range f.folders {
		childList := folder.getFolderSizeList()
		sizeList = append(sizeList, childList...)
	}

	return sizeList
}

func (f *folder) getTotalByLimit(size int) int {
	sum := 0

	if f.size <= size {
		sum += f.size
	}

	for _, folder := range f.folders {
		childSize := folder.getTotalByLimit(size)
		sum += childSize
	}

	return sum
}

func (f *folder) getFolder(dir string) *folder {
	for _, folder := range f.folders {
		if folder.dir == dir {
			return folder
		}
	}

	return nil
}

func (f *folder) addFolder(dir string) {
	newFolder := createNewFolder(dir, f)
	f.folders = append(f.folders, newFolder)
}

func (f *folder) addFile(file *file) {
	f.size += file.size
	f.updateParentSizeValue(file.size)
	f.files = append(f.files, file)
}

func (f *folder) updateParentSizeValue(size int) {
	if f.parentDir != nil {
		f.parentDir.size += size
		f.parentDir.updateParentSizeValue(size)
	}
}

func createNewFolder(dir string, parent *folder) *folder {
	return &folder{
		parentDir: parent,
		dir:       dir,
	}
}

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	return &file{
		name: name,
		size: size,
	}
}

func checkCommand(cmd string, fs *filesystem) {
	values := strings.Split(cmd, " ")
	switch values[0] {
	case "$":
		if values[1] == "cd" {
			fs.goToFolder(values[2])
		}
	case "dir":
		fs.activeFolder.addFolder(values[1])
	default:
		size, err := strconv.Atoi(values[0])
		if err == nil {
			file := newFile(values[1], size)
			fs.activeFolder.addFile(file)
		}

	}
}

func main() {
	file, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fs := createNewFilesystem()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		checkCommand(scanner.Text(), fs)
	}

	fmt.Println("Part 1:", fs.root.getTotalByLimit(MAX_SIZE))

	list := fs.root.getFolderSizeList()
	sort.Ints(list)

	for _, item := range list {
		if (SYSTEM - (fs.root.size - item)) >= UPDATE {
			fmt.Println("Part 2:", item)
			return
		}
	}
}
