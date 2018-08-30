package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// User hoge
type User struct {
	FirstName *string `validate:"required,ne="""`
	LastName  *string `validate:"required,ne="""`
	Age       *int32
	Lmt       *bool `validate:"required"`
}

// SuperUser hoge
type SuperUser struct {
	SuperFirstName *string
	SuperLastName  *string
	SuperAge       *int32
	SuperLmt       *bool
}

var validate *validator.Validate

func ptrStr(s string) *string {
	return &s
}

func ptrInt(s int32) *int32 {
	return &s
}

func ptrBool(i bool) *bool {
	return &i
}

func main() {
	x, err := buildParam()
	if err != nil {
		fmt.Println("errです!")
		fmt.Printf("%+v\n", err)
		return
	}

	//fmt.Println(*x.Lmt)
	fmt.Printf("%#v\n", *x.Lmt)
	ret := &SuperUser{
		SuperFirstName: x.FirstName,
		SuperLastName:  x.LastName,
		SuperAge:       x.Age,
		SuperLmt:       x.Lmt,
	}

	fmt.Printf("%+v\n", ret)

}

func buildParam() (*User, error) {

	validate = validator.New()
	validate.RegisterStructValidation(UserStructLevelValidation, User{})

	u := &User{
		FirstName: ptrStr("5"),
		LastName:  ptrStr("4"),
		//Lmt:       ptrBool(true),
	}

	//err := validate.Struct(u)
	if err := validate.Struct(u); err != nil {
		return nil, err
	}
	//validate.RegisterStructValidation(UserStructLevelValidation, User{})
	fmt.Println("---- valid values")

	return u, nil

}

// UserStructLevelValidation hoge
func UserStructLevelValidation(sl validator.StructLevel) {

	user := sl.Current().Interface().(User)
	if *user.Lmt != true {
		sl.ReportError(*user.Lmt, "lmt", "lmt", "lmt is empty", "")
	}
}
