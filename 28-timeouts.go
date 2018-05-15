package main

// タイムアウトは、外部リソースに接続するプログラムや、実行時間を制限する必要があるプログラムで重要です。
// GOでは、チャネルとselectのおかげでタイムアウトを簡単かるエレガントに実現できます。

// このselectタイムアウトパターンを使うには、結果をチャネル経由でやるとりする必要があります。
// これは一般的によい考えです。というのも、他のGoの重要な機能もチャネルとselectを前提にしているからです。
import (
	"fmt"
	"time"
)

func main() {

	// 例として、2秒後にチャネルc1へ結果を返す外部呼び出しを実行していると仮定しましょう。
	c1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// selectを使ったタイムアウトの実装は次の通りです。
	// res := <-c1が結果を待ち、 <-Time.Afterは1秒のタイムアウト後に送信されてくる値を待ちます。
	// selectは最初に受信したものを処理するので、操作が1秒以上かかるとタイムアウトのケースが選択されます。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// もしタイムアウトをさらに長い3秒にするおｔ、c2からの受信が先に成功し、結果が表示されます。
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}
