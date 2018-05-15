package main

import "fmt"

func f(from string) {

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// このプログラムを実行すると、最初に同期呼び出しの出力、 その次に 2 つのゴルーチンの混ざった出力を確認できます。
	// これは、ゴルーチンが Go ランタイムによって 並行実行されていることを示しています。
	fmt.Scanln()
	fmt.Println("done")
}
