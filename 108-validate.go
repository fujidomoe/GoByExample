package main

import (
	"net/http"

	"fmt"
	"github.com/mholt/binding"
)

type Param struct {
	GroupID *int64
	Version *string
	Lmt     *bool
	Title   *bool
}

// Interface 適合チェック
var _ binding.Validator = (*Param)(nil)

// Validate バリデート
func (p *Param) Validate(r *http.Request) error {
	if p.Title == nil || *p.Title == false {
		return binding.NewError([]string{"title"}, "ValidateError", "title must true")
	}
	return nil
}

// Interface 適合チェック
var _ binding.FieldMapper = (*Param)(nil)

// FieldMap binding
func (p *Param) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&p.GroupID: binding.Field{
			Form:         "G",
			Required:     true,
			ErrorMessage: "G is required",
		},
		&p.Version: "v",
		&p.Lmt: binding.Field{
			Form:         "lmt",
			Required:     true,
			ErrorMessage: "lmt is required",
		},
		&p.Title: "title",
	}
}

// Now your handlers can stay clean and simple
func handler(w http.ResponseWriter, r *http.Request) {
	p := new(Param)
	if errs := binding.Bind(r, p); errs != nil {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		//return
	}
	fmt.Printf("%#v\n", *p)

	//fmt.Fprintf(w, "From:    %d\n", contactForm.User.ID)
	//fmt.Fprintf(w, "Message: %s\n", contactForm.Message)
}

func main() {
	http.HandleFunc("/contact", handler)
	http.ListenAndServe(":3000", nil)
}
