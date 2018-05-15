package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}

	// 複数の引数を素手にスライスで持っている場合は、func(slice...)のような形で可変長引数関数に渡せます。
	sum(nums...)

}
