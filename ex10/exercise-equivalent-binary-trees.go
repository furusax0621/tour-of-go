//Exercise: Equivalent Binary Trees
// 葉に同じ順序の値を保持する、異なる多くの二分木( binary tree )があります。
// 1. Walk 関数を実装してください。
// 2. Walk 関数をテストしてください。
// 3. Same 関数を実装してください。
// 4. Same 関数をテストしてください。
package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// t を走査する関数値(function value)
	var walk func(*tree.Tree)
	walk = func(l *tree.Tree) {
		// 1. t.Left を引数として関数を呼び出す
		if l.Left != nil {
			walk(l.Left)
		}

		// 2. t.Value をチャネルに送る
		ch <- l.Value

		// 3. t.Right を引数として関数を呼び出す
		if l.Right != nil {
			walk(l.Right)
		}
	}

	// t (root of tree)を引数として関数を呼び出す
	walk(t)

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return false
}

func main() {
	// Test Walk
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}
}
