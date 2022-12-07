package day07

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	children NodeMap
	files    *[]string
	size     int
}

type NodeMap map[string]*Node

var RootNode = Node{
	parent:   nil,
	children: make(NodeMap),
	size:     0,
	files:    nil,
}

// Returns 0 for cd, 1 for ls, 2 for dir and 3 for file, string is for payload
func parseLine(line string) (int8, string) {
	matches := strings.Split(line, " ")

	switch len(matches) {
	// 2 groups captured, so either `$ ls`, `dir DIRNAME`, or `SIZE FILENAME`
	case 2:
		if matches[0] == "$" {
			return 1, ""
		}

		if matches[0] == "dir" {
			return 2, matches[1]
		}

		return 3, matches[0] + "," + matches[1] // return size and name of file

	// 3 groups captured, means `$ cd DIRNAME`
	case 3:
		return 0, matches[2]

	default:
		return -1, ""
	}
}

func BuildFileSystemTree(directives []string) Node {
	currentNode := &RootNode

	for _, d := range directives {
		directiveType, payload := parseLine(d)
		if directiveType < 0 {
			panic(1)
		}

		switch directiveType {
		// Means cd
		case 0:
			if payload == "/" {
				currentNode = &RootNode
				continue
			}

			if payload == ".." {
				currentNode = currentNode.parent
				continue
			}

			// if directory has not been discovered yet, add to children
			if !currentNode.children.contains(payload) {
				newNode := Node{
					parent:   currentNode,
					children: make(map[string]*Node),
					size:     0,
					files:    nil,
				}
				currentNode.children[payload] = &newNode
			}

			newNode := currentNode.children[payload]
			currentNode = newNode

		// Means dir listing
		case 2:
			if !currentNode.children.contains(payload) {
				newNode := Node{
					parent:   currentNode,
					children: make(map[string]*Node),
					size:     0,
					files:    nil,
				}
				currentNode.children[payload] = &newNode
			}

		// Means file
		case 3:
			info := strings.Split(payload, ",")
			// if file hasn't been registered before,add to list and add size to dir's size
			if currentNode.files == nil {
				newList := make([]string, 0)
				currentNode.files = &newList
			}
			if !contains(info[1], currentNode.files) {
				newList := append(*currentNode.files, info[1])
				currentNode.files = &newList
				size, err := strconv.ParseInt(info[0], 10, 64)
				if err != nil {
					fmt.Print(err)
					panic(1)
				}
				currentNode.size += int(size)
			}
		default:
		}
	}
	return RootNode
}

func (n *NodeMap) contains(node string) bool {
	for k := range *n {
		if k == node {
			return true
		}
	}
	return false
}

func contains(file string, f *[]string) bool {
	for _, k := range *f {
		if k == file {
			return true
		}
	}
	return false
}

func (n *Node) GetSize() int {
	size := n.size
	if len(n.children) > 0 {
		for _, child := range n.children {
			size += child.GetSize()
		}
	}

	return size
}

func (n *Node) GetSizeBelowNumber(number int) int {
	if len(n.children) == 0 {
		if n.size > number {
			return 0
		} else {
			return n.size
		}
	}

	downstreamSize := 0
	for _, c := range n.children {
		downstreamSize += c.GetSizeBelowNumber(number)
	}

	ownSize := n.GetSize()
	if ownSize > number {
		return downstreamSize
	} else {
		return downstreamSize + ownSize
	}
}

func (n *Node) GetSmallestAbove(above int, hist *int) {
	treeSize := n.GetSize()

	if treeSize < *hist && treeSize >= above {
		*hist = treeSize
	}

	if n.children != nil && len(n.children) != 0 {
		for k := range n.children {
			n.children[k].GetSmallestAbove(above, hist)
		}
	}
}
