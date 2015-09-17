package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, level int) {
	if t.Left != nil {
		Walk(t.Left, ch, level+1)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch, level+1)
	}

	if level == 0 {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1, 0)
	go Walk(t2, c2, 0)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		// fmt.Println(v1)
		// fmt.Println(v2)
		if (ok1 && !ok2) || (!ok1 && ok2) {
			return false
		}

		if !ok1 && !ok2 {
			return true
		}

		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
