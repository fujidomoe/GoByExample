package main

import (
	"fmt"
	"sort"
)

// Goでカスタム関数を使ってソートする為には対応する型が必要です。
// ここではbyLength型をつくりました。
// これは、組み込みの[]string型のただのエイリアスです。
type byLength []string

// sortパッケージのSort関数を使える用に、sort.Interfaceすなわり、Len, less,Sqwap関数を実装します。
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	// 元のfruitsスライスをbyLengthにキャストし、sort.Sort関数を使うことでカスタムソートを実現できます。
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
