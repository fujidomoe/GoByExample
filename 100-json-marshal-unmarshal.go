package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Job  string `json:"job,omitempty"`
}

func main() {
	u := &User{}
	u.Name = "haru"
	u.Age = 4

	fmt.Println(u) // &{haru 4 }

	// struct to json
	bytes, _ := json.Marshal(&u)
	fmt.Println(string(bytes)) // {"name":"haru","age":4}

	// json to struct
	data := &User{}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		panic(err)
	}

	data.Job = "Kindergarten child"
	fmt.Println(data) // &{haru 4 Kindergarten child}

	bytes, _ = json.Marshal(data)
	fmt.Println(string(bytes)) // {"name":"haru","age":4,"job":"Kindergarten child"}
}
