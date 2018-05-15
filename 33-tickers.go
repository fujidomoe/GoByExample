package main

import (
	"fmt"
	"time"
)

// タイマーは未来に一度だけしたい時に使いますが、ティッカーは定期的に何かしたい時に使います。
// ここでは、停止するまで定期的に動作すティッカーの例を見ます。

func main() {

	// ティッカーはタイマーと同様の仕組み、すなわち値を送信するチャネルを使います。
	// ここでは、500ミリ秒毎に受信される値を繰り返し処理する為に、rangeビルトイン関数を使っています。
	ticker := time.NewTicker(500 * time.Microsecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// ティッカーはタイマーと同じ用に停止できます。
	// ティッカーが停止されると、そのチャネルからはも値を受信しなくなります。
	// この例では 1600ミリ秒後に停止します。
	time.Sleep(1600 * time.Microsecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	// このプログラムを実行すると、ティッカーは停止されるまでに3回動作すいるはずです。
}
