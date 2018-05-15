package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p) // {1 2}

	fmt.Printf("%+v\n", p) // {x:1 y:2}

	fmt.Printf("%#v\n", p) // main.point{x:1, y:2}

}
