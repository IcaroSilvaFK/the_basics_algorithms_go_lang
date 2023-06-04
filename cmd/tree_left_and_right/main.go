package main

import (
	"fmt"
	"time"
)

type Tree struct {
	Name  string
	Time  time.Time
	left  *Tree
	right *Tree
}

func main() {

	first := &Tree{
		Name:  "first",
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}

	second := &Tree{
		Name:  "second",
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	third := &Tree{
		Name:  "third",
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	fourth := &Tree{
		Name:  "fourth",
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	five := &Tree{
		Name:  "five",
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}

	left := &Tree{
		Name:  "First left",
		Time:  time.Now(),
		left:  nil,
		right: nil,
	}
	right := &Tree{
		Name:  "First left right",
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
	findNameOnTree(first, "First left right")

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
		fmt.Println(node.Name)
	}
}

func findNameOnTree(t *Tree, name string) {

	stack := []*Tree{t}
	iterations := 0

	for len(stack) > 0 {

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Name == name {
			fmt.Println(iterations)
			fmt.Println(node.Name)
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
