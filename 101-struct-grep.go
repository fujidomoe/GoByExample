package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Ext struct {
	Ads []*Ad
}
type ExtV []*Ad

func (a ExtV) Len() int           { return len(a) }
func (a ExtV) Less(i, j int) bool { return *a[i].Priority < *a[j].Priority }
func (a ExtV) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Ad struct {
	VendorName  *string  `json:"vendor_name"`
	Priority    *int32   `json:"priority"`
	Imptrackers []string `json:"imptrackers"`
}

func strPtr(s string) *string {
	return &s
}

func intPtr(i int32) *int32 {
	return &i
}

func main() {
	ad1 := &Ad{
		VendorName:  strPtr("softbank"),
		Priority:    intPtr(20),
		Imptrackers: []string{"hoge"},
	}
	ad2 := &Ad{
		VendorName:  strPtr("docomo"),
		Priority:    intPtr(40),
		Imptrackers: []string{"hoge"},
	}
	ad3 := &Ad{
		VendorName:  strPtr("au"),
		Priority:    intPtr(30),
		Imptrackers: []string{"hoge"},
	}

	ext := &Ext{
		Ads: []*Ad{ad1, ad2, ad3},
	}

	//ads := []*Ad{ad1, ad2, ad3}

	x := ext.getOne()
	//ad := &Ad{
	//	VendorName:  x.VendorName,
	//	Priority:    x.Priority,
	//	Imptrackers: x.Imptrackers,
	//}
	//r := make([]string, 1)
	//s := []string{}
	//s = append(s, "fuga")
	//s = append(s, x.Imptrackers)
	x.Imptrackers = append(x.Imptrackers, "fuga")

	//r = append(r, "fuga", x.Imptrackers)
	y := &Ext{
		Ads: []*Ad{
			{
				VendorName:  x.VendorName,
				Priority:    x.Priority,
				Imptrackers: x.Imptrackers,
			},
		},
	}
	bytes, _ := json.Marshal(y)
	fmt.Printf("%s", bytes)
}

// struct配列に対して、Priorityで昇順ソートして、0番目のstructにPriorityを振り直して返す
//func (x *Ext) getOne() *ExtV {
//	one := make(ExtV, 0)
//	sort.Sort(ExtV(x.Ads))
//	x.Ads[0].Priority = intPtr(0)
//	one = append(one, x.Ads[1])
//	return &one
//}

func (x *Ext) getOne() *Ad {
	//one := make(ExtV, 0)
	sort.Sort(ExtV(x.Ads))
	x.Ads[0].Priority = intPtr(0)
	//one = append(one, x.Ads[1])
	//return &one
	return x.Ads[0]
}
