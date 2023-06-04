package main

import (
	"fmt"
	"time"
)

type Tree struct {
	Name  string
	Id    int
	Time  time.Time
	left  *Tree
	right *Tree
}

func main() {

	first := &Tree{
		Name:  "first",
		Id:    1,
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}

	second := &Tree{
		Name:  "second",
		Id:    2,
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	third := &Tree{
		Name:  "third",
		Id:    3,
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	fourth := &Tree{
		Name:  "fourth",
		Id:    4,
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	five := &Tree{
		Name:  "five",
		Id:    5,
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}

	left := &Tree{
		Name:  "left",
		Id:    6,
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	right := &Tree{
		Name:  "right",
		Id:    7,
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	first.right = second
	first.left = left
	second.right = third
	second.left = right
	left.right = fourth
	fourth.right = five

	/*

	                         first
	   											/  \
	   									second  left
	   									/   \     \
	   							third  right  fourth
	   							               |
	   														five

	*/

	walkOnTree(first)
	findNameOnTree(first, "right")
	finded := findNameOnTreeRecursive(first, "right")

	fmt.Println("Finded:", finded)

	result := walkOnTreeRecursive(first)

	for _, v := range result {
		fmt.Println(*v)
	}

	breadthFirstValues(first)
}

func walkOnTree(t *Tree) {

	stack := []*Tree{t}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.left != nil {
			stack = append(stack, node.left)
		}
		if node.right != nil {
			stack = append(stack, node.right)
		}
		fmt.Printf("Node name:%s . Node id: %d\n", node.Name, node.Id)
	}
}

func walkOnTreeRecursive(t *Tree) []*Tree {
	stack := []*Tree{t}

	if t == nil {
		return nil
	}

	left := walkOnTreeRecursive(t.left)
	right := walkOnTreeRecursive(t.right)

	// stack = append(stack, walkOnTreeRecursive(t.left)...)
	// stack = append(stack, walkOnTreeRecursive(t.right)...)

	stack = append(stack, left...)
	stack = append(stack, right...)

	return stack
}

func findNameOnTree(t *Tree, name string) {

	stack := []*Tree{t}
	iterations := 0

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Name == name {
			fmt.Printf("Iterations: %d\nNode name: %s . Node id: %d\n", iterations, node.Name, node.Id)
			return
		}

		if node.left != nil {
			stack = append(stack, node.left)
		}
		if node.right != nil {
			stack = append(stack, node.right)
		}
		iterations++
	}
}

func findNameOnTreeRecursive(t *Tree, name string) bool {

	if t.Name == name {
		fmt.Println(" Node name:", t.Name, "Node id:", t.Id)
		return true
	}
	if t.left == nil {
		return false
	}
	if t.right == nil {
		return false
	}

	return findNameOnTreeRecursive(t.right, name) || findNameOnTreeRecursive(t.left, name)
}

func breadthFirstValues(t *Tree) {

	stack := []*Tree{t}
	walk := []string{}

	for len(stack) > 0 {

		node := stack[0]
		walk = append(walk, node.Name)
		stack = stack[1:]

		if node.right != nil {
			stack = append(stack, node.right)
		}
		if node.left != nil {
			stack = append(stack, node.left)
		}

	}

	fmt.Println(walk)

}
