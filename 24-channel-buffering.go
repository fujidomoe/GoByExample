package main

import "fmt"

// デフォルトではチャネルはバッファリングされないので、受信(<- chan)の準備ができていないと送信(chan <-)できません
// しかし、バッファリングされたチャネルは、対応する受信ががいなくても決められた量までなら値を送信することができます。

func main() {

	// この例ではstringを2つまでバッファリングするチャネルをmakeしています。
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
