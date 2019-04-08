// Exercise: Maps
// WordCount 関数を実装してみましょう。string s で渡される文章の、各単語の出現回数のmapを返す必要があります。
// wc.Test 関数は、引数に渡した関数に対しテストスイートを実行し、成功か失敗かを結果に表示します。
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount 関数は、string s で渡される文章の、各単語の出現回数のmapを返します。
func WordCount(s string) map[string]int {
	wcmap := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		_, ok := wcmap[word]
		if ok {
			wcmap[word]++
		} else {
			wcmap[word] = 1
		}
	}

	return wcmap
}

func main() {
	wc.Test(WordCount)
}
