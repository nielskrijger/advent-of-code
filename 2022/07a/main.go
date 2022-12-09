package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var minFileSize = 100_000

func main() {
	data, _ := os.ReadFile("2022/07a/sample.txt")
	tree := createFileTree(strings.Split(string(data), "\n"))

	fmt.Printf("Answer: %d", findCandidatesFileSize(tree))
}

type file struct {
	name   string
	size   int
	isDir  bool
	parent *file
	files  []*file
}

func (f *file) addFile(name string, size int) {
	f.files = append(f.files, &file{name: name, parent: f})
	f.increaseSize(size)
}

func (f *file) addDir(name string) {
	f.files = append(f.files, &file{name: name, parent: f, files: []*file{}, isDir: true})
}

func (f *file) increaseSize(size int) {
	f.size += size
	if f.parent != nil {
		f.parent.increaseSize(size)
	}
}

func (f *file) findDir(name string) *file {
	for _, d := range f.files {
		if d.name == name && d.isDir {
			return d
		}
	}
	return nil
}

func findCandidatesFileSize(dir *file) int {
	total := 0

	if dir.size <= minFileSize {
		total += dir.size
	}

	for _, file := range dir.files {
		if file.isDir {
			total += findCandidatesFileSize(file)
		}
	}

	return total
}

func createFileTree(lines []string) *file {
	root := &file{name: "/", files: []*file{}}
	current := root

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		switch {
		case line == "$ cd /":
			current = root
		case line == "$ cd ..":
			current = current.parent
		case strings.HasPrefix(line, "$ cd "):
			pieces := strings.Split(line, " ")
			current = current.findDir(pieces[2])
		case line == "$ ls":
			for {
				// Peek, stop if we find another command or reached end of input
				if i == len(lines)-1 || lines[i+1][0] == '$' {
					break
				}

				// Consume next line
				i++
				line = lines[i]
				pieces := strings.Split(line, " ")

				// Add file or dir
				if pieces[0] == "dir" {
					current.addDir(pieces[1])
				} else {
					size, _ := strconv.Atoi(pieces[0])
					current.addFile(pieces[1], size)
				}
			}
		}
	}

	return root
}
