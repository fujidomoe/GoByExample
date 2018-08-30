package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

type User struct {
	Name *string `validate:required`
	Age  *int32  `validate:"min=5"`
	Lmt  *bool
}

func ptrStr(s string) *string {
	return &s
}

func ptrInt(i int32) *int32 {
	return &i
}

func main() {
	x, err := buildParam()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", x)
}

func buildParam() (*User, error) {
	validate = validator.New()

	u := &User{
		Name: ptrStr("haru"),
		Age:  ptrInt(4),
	}

	/*
		if err := v.Struct(u); err == nil {
			fmt.Println("Valid!")
		} else {
			fmt.Println("Invalid!", err)

		}
	*/
	if err := v.Struct(u); err != nil {
		return nil, err
	}
	fmt.Println("---- valid values")

	return u, nil

}
