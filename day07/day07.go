package day07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	children NodeMap
	files    Files
	size     int
}

type NodeMap map[string]Node
type Files []string

var rootNode = Node{
	parent:   nil,
	children: make(NodeMap),
	size:     0,
	files:    make(Files, 0),
}

// Returns 0 for cd, 1 for ls, 2 for dir and 3 for file, string is for payload
func parseLine(line string) (int8, string) {
	regexp := regexp.MustCompile(`(\S+) (\S+) ?(\S+)?`)
	matches := regexp.FindStringSubmatch(line)

	switch len(matches) {
	// 2 groups captured, so either `$ ls`, `dir DIRNAME`, or `SIZE FILENAME`
	case 3:
		if matches[1] == "$" {
			return 1, ""
		}

		if matches[1] == "dir" {
			return 2, matches[2]
		}

		return 3, matches[1] + "," + matches[2] // return size and name of file

	// 3 groups captured, means `$ cd DIRNAME`
	case 4:
		return 0, matches[3]

	default:
		return -1, ""
	}
}

func BuildFileSystemTree(directives []string) int8 {
	currentNode := rootNode

	for _, d := range directives {
		directiveType, payload := parseLine(d)
		if directiveType < 0 {
			return directiveType
		}

		switch directiveType {
		// Means cd
		case 0:
			if payload == "/" {
				currentNode = rootNode
			}

			if payload == ".." {
				currentNode = *currentNode.parent
			}

			// if directory has not been discovered yet, add to children
			if !currentNode.children.contains(payload) {
				currentNode.children[payload] = Node{
					parent:   &currentNode,
					children: make(map[string]Node),
					size:     0,
				}
			}
			currentNode = currentNode.children[payload]

		// Means dir listing
		case 2:
			if !currentNode.children.contains(payload) {
				currentNode.children[payload] = Node{
					parent:   &currentNode,
					children: make(map[string]Node),
					size:     0,
				}
			}

		// Means file
		case 3:
			info := strings.Split(payload, ",")
			// if file hasn't been registered before,add to list and add size to dir's size
			if !currentNode.files.contains(info[1]) {
				currentNode.files = append(currentNode.files, info[1])
				size, err := strconv.ParseInt(info[0], 10, 64)
				if err != nil {
					fmt.Print(err)
					return -1
				}

				currentNode.size += int(size)
			}
		default:
		}
	}
	return 0
}

func (n *NodeMap) contains(node string) bool {
	for k := range *n {
		if k == node {
			return true
		}
	}
	return false
}

func (f *Files) contains(file string) bool {
	for _, k := range *f {
		if k == file {
			return true
		}
	}
	return false
}

func (n *Node) getSize() int {
	size := n.size
	if len(n.children) > 0 {
		for _, child := range n.children {
			size += child.getSize()
		}
	}

	return size
}
