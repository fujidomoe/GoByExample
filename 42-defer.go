package main

// deferは通常クリーンアップなどの為に、関数呼び出しが後で実行される事を保証する為に使います。
// deferは、他のぷrp具ラミング言語でensureやfinallyが使われる場面でよく使われます。
import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("close")
	f.Close()
}
