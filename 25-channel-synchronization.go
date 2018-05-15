package main

// ゴルーチン間の実行を同期する為に、チャネルを使う事ができます。
// ここでの例は、ゴルーチンの完了を待つためにブロッキング受信を使います。

import (
	"fmt"
	"time"
)

// ゴルーチンで実行する関数はこちらです。
// この関数が完了したことを別のゴルーチンに通知ｓるう為に、doneチャネルが使われます。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true // 完了した事を通知する為に値を送信します。
}

func main() {

	// 通知陽のチャネルを私て、workerゴルーチンを開始します。
	done := make(chan bool, 1)
	go worker(done)
	<-done // チャネルへ完了通知を受信するまでブロックします。

}
