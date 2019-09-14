package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {

	defer close(ch)

	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, o1 := <-c1
		v2, o2 := <-c2
		fmt.Printf("%d : %d\n", v1, v2)
		if !(v1 == v2 && o1 == o2) {
			return false
		} else if !o1 {
			return true
		}
	}

}

func main() {

	t1 := tree.New(30)
	t2 := tree.New(30)
	fmt.Printf("t1 and t2 is same: %v\n", Same(t1, t2))
}

