package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	fmt.Println(n)
	if len(parts) == height {
		n.pattern = pattern
		fmt.Println(n)
		fmt.Println("add node pattern:",n.pattern)
		return
	}

	fmt.Println("height:",height)
	fmt.Println("insert parts:",parts)
	part := parts[height]

	fmt.Println("part:",part)
	child := n.matchChild(part)
	fmt.Println("children",n.children)
	fmt.Println("child",child)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		fmt.Println(child)
		n.children = append(n.children, child)
	}
	fmt.Println("children",n.children)
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	fmt.Println("search parts:",parts)
	fmt.Println("n.part:",n.part)
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	fmt.Println("n",n)
	fmt.Println("match before children:",n.children)
	children := n.matchChildren(part)
	fmt.Println("children:",children)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}

func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}

func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

