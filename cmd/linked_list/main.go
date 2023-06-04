package main

import (
	"fmt"
	"time"
)

type Tree struct {
	Name string
	Time time.Time
	left *Tree
}

func main() {

	first := &Tree{
		Name: "first",
		Time: time.Now(),
		left: nil,
	}
	second := &Tree{
		Name: "second",
		Time: time.Now(),
		left: first,
	}
	third := &Tree{
		Name: "third",
		Time: time.Now(),
		left: second,
	}

	walkOnThree(third)

	findInLinkedList(third, "second")

}

func walkOnThree(t *Tree) {
	if t.left != nil {
		fmt.Println(*t)
		walkOnThree(t.left)
	} else {
		fmt.Println("Last child of the tree")
		fmt.Println(*t)
		fmt.Println("End of the tree")
	}
}

func findInLinkedList(t *Tree, name string) {

	if t.Name == name {
		fmt.Println(*t)
		fmt.Println("Found")
		return
	}

	if t.left != nil {
		findInLinkedList(t.left, name)
	}
}
