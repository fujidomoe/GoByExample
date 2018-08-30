package main

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
)

// Param Sample
type Param struct {
	MIMES *[]string `schema:"mimes"`
	IP    *string   `schema:"ip"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	param := &Param{}
	u := r.URL.Query()
	u.Add("ip", "200.200.200")
	d := schema.NewDecoder()
	d.Decode(param, u)
	fmt.Printf("%#v\n", *param.MIMES)

	for _, v := range *param.MIMES {
		fmt.Println(v)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
