package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Bob", age: 20})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40}) // 構造体へのポインタになります

	s := person{name: "Sean", age: 40}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // 構造体へのポインタにもドットが使えます。この場合、ポインタは自動的にデリファレンスされます

	sp.age = 51
	fmt.Println(sp.age)

}
