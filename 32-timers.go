package main

// 将来のある時点や一定間隔で繰り返しGoコードを十個うする時に利用します。
// 今回はタイマーを見ていきます。

import (
	"fmt"
	"time"
)

func main() {
	// タイマーは、将来の1回限りのイベントを表します。
	// タイマーに待ち時間を指定すると、指定した時間が経過した後で通知してくれるチャネルが提供されます。
	// このタイマーは2秒間まちます。
	timer1 := time.NewTicker(2 * time.Second)

	//　timer1.CはタイマーのチャネルCが時間になったことを知らせる値を送信するまでブロックします。
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// ただ待ちたいだけれあれば、time.Sleepが使えます。
	// タイマーが役に立つ1つの理由は指定時間が経過する前にキャンセルすることができる点です。
	// これがその例です。
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 1つ目のタイマーは、プログラムを開始してから2秒後には完了しますが、
	// 2つ目のタイマーは完了する前に停止させられるはずです。
}
