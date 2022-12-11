package main

import (
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Dir struct {
	Name   string
	Dirs   []*Dir
	Files  []*File
	Parent *Dir
}

func (dir *Dir) Size() int {
	total := 0
	for _, file := range dir.Files {
		total += file.Size
	}
	for _, dir := range dir.Dirs {
		total += dir.Size()
	}
	return total
}

func (d *Dir) Key() string {
	if d.Parent == nil {
		return d.Name
	}
	return d.Parent.Key() + d.Name + "/"
}

func parse(file string) Dir {
	lines := strings.Split(file, "\n")
	pointers := map[string]*Dir{}
	current := "/"
	pointers[current] = &Dir{"/", []*Dir{}, []*File{}, nil}
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "$ ls") {
			continue
		} else if strings.HasPrefix(line, "$ cd ..") {
			if p, ok := pointers[current]; ok {
				current = p.Parent.Key()
			}
		} else if strings.HasPrefix(line, "$ cd") {
			segments := strings.Split(line, " ")
			if p, ok := pointers[current + segments[2] + "/"]; ok {
				current = p.Key()
			}
		} else if strings.HasPrefix(line, "dir") {
			segments := strings.Split(line, " ")
			if p, ok := pointers[current]; ok {
				child := Dir{segments[1], []*Dir{}, []*File{}, p}
				pointers[child.Key()] = &child
				p.Dirs = append(p.Dirs, &child)
			}
		} else {
			segments := strings.Split(line, " ")
			num, _ := strconv.Atoi(segments[0])
			if p, ok := pointers[current]; ok {
				file := File{segments[1], num}
				p.Files = append(p.Files, &file)
			}
		}
	}
	return *pointers["/"]
}

func sizes(dirs []*Dir, total []int) []int {
	if len(dirs) == 0 {
		return total
	}
	dir := dirs[0]
	total = append(total, dir.Size())
	rest := dirs[1:]
	rest = append(rest, dir.Dirs...)
	return sizes(rest, total)
}

func day07(file string) int {
	dir := parse(file)
	dirSizes := sizes([]*Dir{&dir}, []int{})
	total := 0
	for _, size := range dirSizes {
		if size <= 100000 {
			total += size
		}
	}
	return total
}
