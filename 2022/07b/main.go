package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const totalDiskSize = 70_000_000
const requiredSpace = 30_000_000

func main() {
	data, _ := os.ReadFile("2022/07b/input.txt")
	tree := createFileTree(strings.Split(string(data), "\n"))

	// Find all directories that would free up enough space
	unusedSpace := totalDiskSize - tree.size
	minDirSpace := requiredSpace - unusedSpace
	candidates := findCandidates(tree, minDirSpace)

	// Sort from min size to max size
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].size < candidates[j].size
	})

	fmt.Printf("Answer: %v", candidates[0].size)
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

func findCandidates(directory *file, minSize int) []*file {
	dirs := make([]*file, 0)

	if directory.size >= minSize { // this is the size we need to free up
		dirs = append(dirs, directory)
	}

	for _, subdir := range directory.files {
		dirs = append(dirs, findCandidates(subdir, minSize)...)
	}

	return dirs
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
				// Peek next line, stop if we find a command or reached end of input
				if i == len(lines)-1 || lines[i+1][0] == '$' {
					break
				}

				// Consume next line
				i++
				line = lines[i]
				pieces := strings.Split(line, " ")

				// Add file
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
