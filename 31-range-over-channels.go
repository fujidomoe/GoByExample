package main

// 11-range.goの例で、forとrangeが基本的なデータ構造に対して、どのように反復処理を提供するかを見ました。
// この構文は、チャネルから受信した値を反復処理する場合にも使うことができます。

// 今回の例は、受信すべき値がまだ残っている、空でないチャネルでもクローズできることを示しています。
import "fmt"

func main() {
	// queueチャネルの2つの値を反復処理するとします。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// このrangeは、queueから受信した要素を反復処理します。
	// 上でチャネルをcloseしたので、反復処理は2つの要素を受信した時に終了します。
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
