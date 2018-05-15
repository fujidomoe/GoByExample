package main

// チャネル上の基本的な送受信はブロックします。
// しかし、default句をもったselectを使えば、ブロックしない(non-blocking)送信、受信、そして多重selectさえ実装することができます
import "fmt"

func main() {

	messages := make(chan string)
	signals := make(chan bool)

	// これは、ブロックしない受信の例です。
	// もしmessagesチャネルで値が準備できていれは、selectは<-messagesのケースを実行します。
	// そうでなければ、すぐにdefaultのケースを実行します。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// ブロックしない送信も同様にど動作します。
	// messagesチャネルには、バッファもなければ受ける側も無いため、msgを送信することはできません。
	// そのため、defaultケースが実行されます。
	msg := "hi"

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// ブロックしない多重selectを実現する為に、default句の前に複数のcaseを使うことができます。
	// 次の例では、messagesとsignalsの両方に対してブロックしない受信をこころみます。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
