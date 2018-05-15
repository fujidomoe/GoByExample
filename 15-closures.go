package main

import "fmt"

func intSeq() func() int {
	fmt.Println("called")
	i := 0
	return func() int {
		i++
		return i
	}
}
func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 関数ごとに個別の状態を持っている事を確認する為にもう1つ新たに作成して試してみる
	newInts := intSeq()
	fmt.Println(newInts())

}
