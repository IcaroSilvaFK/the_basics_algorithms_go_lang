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
	tree := &Tree{
		Name: "Firs",
		Time: time.Now(),
		left: &Tree{
			Name: "First left",
			Time: time.Now(),
			left: &Tree{
				Name: "First left left",
				Time: time.Now(),
				left: nil,
				right: &Tree{
					Name: "First left right",
					Time: time.Now(),
					left: &Tree{
						Name:  "First left right left",
						Time:  time.Now(),
						left:  nil,
						right: nil,
					},
					right: nil,
				},
			},
			right: &Tree{
				Name: "First left right",
				Time: time.Now(),
				left: nil,
				right: &Tree{
					Name: "First left right right",
					Time: time.Now(),
					left: &Tree{
						Name: "First left right right left",
						Time: time.Now(),
						left: nil,
						right: &Tree{
							Name:  "First left right right right",
							Time:  time.Now(),
							left:  nil,
							right: nil,
						},
					},
					right: nil,
				},
			},
		},
	}

	walkOnThree(tree)

}

func walkOnThree(t *Tree) {

	if t.left != nil {
		fmt.Println("Left")
		walkOnThree(t.left)
		fmt.Println(*t)
	}
	if t.right != nil {
		fmt.Println("Right")
		walkOnThree(t.right)
		fmt.Println(*t)
	}

}
