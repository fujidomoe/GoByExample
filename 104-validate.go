package main

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// User contains user information
type User struct {
	Name *string `validate:"required"`
	Age  *int32  `validate:"gte=0,lte=130"`
	//.Lmt  *bool   `validate:"required,eq=true"`
	Lmt *bool `validate:"eq=true"`
}

func ptrStr(s string) *string {
	return &s
}

func ptrInt(i int32) *int32 {
	return &i
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// hoge
func main() {
	validate = validator.New()
	user := validateStruct()
	fmt.Printf("%#v\n", *user.Age)
}

func validateStruct() *User {

	user := &User{
		Name: ptrStr("Smith"),
		Age:  ptrInt(13),
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return user
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return user
	}
	return user
}
