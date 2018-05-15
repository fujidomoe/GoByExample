package main

// Goで状態を管理する為の最も重要な仕組みは、チャネルを使った通信です。
// これについては`34-worker-pools`の例でみました。
// しかし、状態を管理する方法はほかにもいくつかあります。
// ここでは、複数のゴルーチンからアクセスされるアトミックカウンターの為のsync/atomicパッケージを使う方法を見ていきます。
import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// カウンターのために符号なし整数を使います。
	var ops uint64
	// 同時更新をシュミレートする為に、1ミリ秒に1回カウンターをインクリメントするゴルーチンを50個開始します。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				//カウントをアトミックにインクリメントする為、&構文でopsカウンターのメモリアドレスをAddUnit64に与えます。
				atomic.AddUint64(&ops, 1)
				// 次のインクリメントまで少し待つ
				time.Sleep(time.Microsecond)
			}
		}()
	}

	// 加算が進むように1秒待ちます。
	time.Sleep(time.Second)

	// まだ他のゴルーチンで更新され続けているカウンターを安全に使う為、
	// LoadUnit64で現在の値のコピーをopsFinalに取り出します。
	// 先の例と同様に、この関数にも&opsで値の取得元のメモリアドレスを与える必要があります。
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
