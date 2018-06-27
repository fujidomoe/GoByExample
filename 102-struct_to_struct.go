package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name,omitempty"`
	Age  string `json:"age,omitempty"`
	//Age string `json:"age"`
}

type User struct {
	Name string `json:"name,omitempty"`
	Age  string `json:"age,omitempty"`
}

func main() {

	p := &Person{Name: "sakura1116"}
	bytes, _ := json.Marshal(&p)
	fmt.Println(string(bytes))
	data := &Person{}
	json.Unmarshal(bytes, data)
	fmt.Println(data)
	u := &User{
		Name: data.Name,
		Age:  data.Age,
	}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
}
