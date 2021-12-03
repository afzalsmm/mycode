package binsrch

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}
type bst struct {
	root *node
	len  int
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}
func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}
func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	b.inOrderTraversalByNode(sb, root.right)
}

/*
if root is nil, add vale as the root node
if node is nil, add value as node
Else else call add on left or right
*/
func (b *bst) addNode(root *node, value int) *node {
	if root == nil {
		b.len++
		return &node{value: value}
	}
	if value < root.value {
		root.left = b.addNode(root.left, value)
	} else if value > root.value {
		root.right = b.addNode(root.right, value)
	}
	return root
}
func (b *bst) add(value int) {
	b.root = b.addNode(b.root, value)
}

func (b bst) searchNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}
	if value < root.value {
		return b.searchNode(root.left, value)
	} else if value > root.value {
		return b.searchNode(root.right, value)
	} else {
		return root, true
	}
}
func (b bst) search(value int) (*node, bool) {
	return b.searchNode(b.root, value)
}

func createTree() *bst {
	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{3, nil, nil}
	b := bst{
		root: n,
		len:  3,
	}

	//b = bst{root: nil, len: 0}
	fmt.Println(b)
	return &b
}

func addNodes() *bst {
	b := createTree()

	b.add(2)
	b.add(5)
	b.add(10)
	fmt.Println(b, &b)
	return b
}

func search() {
	b := addNodes()

	fmt.Println(b.search(1))
	fmt.Println(b.search(0))
	fmt.Println(b.search(2))
	fmt.Println(b.search(6))
	fmt.Println(b.search(4))
	fmt.Println(b.search(10))
}

