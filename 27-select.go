package main

// Goのselectを使うと、複数のチャネル操作を待つことができます。
// ゴルーチンとチャネルをselectで扱えるのがGoの協力な特徴です。
import (
	"fmt"
	"time"
)

func main() {

	// 2つのチャネルに対してselectする例を見ていきます。
	c1 := make(chan string)
	c2 := make(chan string)

	// 各チャネルは言って時間語に値を受信します。
	// これは、例えば同期的なRPC操作をゴルーチンで並行実行する場合をシュミレートしています。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// これらの値を同時に待つためにselectを使い受信したものから画面に表示します。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

}
