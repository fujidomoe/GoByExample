package main

import "os"

func main() {

	// 予期せぬエラーを確認する為に、panicを使います。
	// これはpanicするために作られた唯一のプログラムです。
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
