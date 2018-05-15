package main

// チャネルを関数の引数として使う場合、そのチャネルが送信または、受信だけを意図しているか指定することができます
// これは、プログラムをより型安全にします。

import "fmt"

// このping関数は、チャネルを送信専用で受取ります。
// このチャネルで受信しようとすると、コンパイルエラーになります。

func ping(pings chan<- string, msg string) {
	pings <- msg
}

// このpong関数は、1つ目のチャネルを受信専用で(pings)、2つ目のチャネルを送信専用で(pongs)受取ます。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
