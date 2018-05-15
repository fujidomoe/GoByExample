package main

// チャネルをクローズする(closing)ころは、もう値を送信しないことを意味します。
// これは、チャネルの受けてに完了を伝えるのに便利です。

import "fmt"

func main() {

	// この例では、main()ゴルーチンからワーカーのゴルーチンへタスクの完了を伝える為に、
	// jobsチャネルおｗ使います。ワーカータスクがなくなれば、jobsチャネルをcloseします。
	jobs := make(chan int, 5)
	done := make(chan bool)

	// ワーカーのゴルーチンは次の通りです。
	// `j, more := <-jobs`でjobsチャネルから繰り返し受信します。
	// この2値の形式の受信では、jobsがcloseされ、チャネルの全ての値がすでに受信されていれば、moreの値がfalseになります。
	// ここでは、全てのタスクが完了した時にdoneチャネルへ通知する為に使っています。
	go func() {
		for {
			j, more := <-jobs
			fmt.Println(more)
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// これは、jobsチャネルを通して3つのジョブをワーカーへ送信し、その後チャネルをクローズします。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done // すでに学んだ チャネルの同期手法を使って、ワーカーの完了を待ちます。
	fmt.Println("fin..")
}
