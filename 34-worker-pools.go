package main

// ここではゴルーチンとチャネルを使って、ワーカープール(worker pool)を実装する例をみていきます。
import (
	"fmt"
	"time"
)

// これは、複数のインスタううすが並行実行されるワーカーです。
// これらのワーカーは、jobsチャネルからタスクを受信し、結果をresultsチャネルに送信します。
// 実行コストの高いジョブをシュミレートする為、各タスクは1秒スリープします。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// ワーカープールを使う為には、タスクワーカーに粗送信し、それらの結果を集まる必要があります。
	// その為に、2つのチャネルを作成します。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// ここで、3つのワーカーを開始しますが、最初はまだジョブが無いためブロックします。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 次に、5つのジョブを送信し、それが全てであることを伝える為にチャネルをcloseします。
	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	close(jobs)
	for a := 1; a <= 5; a++ {
		fmt.Println("###", a)
		<-results // 最後に、全てのタスクの結果を集めます。
	}
}

// 実行したプログラムは、5つのジョブが様々なワーカーで実行されていることを示します。
// このプログラムは合計で5秒分のタスクを実行しているにもかかわらず、2秒しかかかりません。
// というのも、3つのワーカーが並行実行されるからです。
