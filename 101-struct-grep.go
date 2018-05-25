package main

import (
	"encoding/json"
	"fmt"
)

type Ext struct {
	Ads []*Ad
}
type ExtV []Ad

type Ad struct {
	VendorName *string `json:"vendor_name"`
	Priority   *int32  `json:"priority"`
}

func strPtr(s string) *string {
	return &s
}

func intPtr(i int32) *int32 {
	return &i
}

func main() {

	ad1 := &Ad{
		VendorName: strPtr("softbank"),
		Priority:   intPtr(1),
	}
	ad2 := &Ad{
		VendorName: strPtr("docomo"),
		Priority:   intPtr(2),
	}

	ext := &Ext{
		Ads: []*Ad{ad1, ad2},
	}

	x := ext.Filter(intPtr(2))

	bytes, _ := json.Marshal(x)
	fmt.Printf("%s", bytes)
}

func (ext *Ext) Filter(i *int32) ExtV {
	filtered := make(ExtV, 0)
	for _, ad := range ext.Ads {
		if *ad.Priority == *i {
			filtered = append(filtered, *ad)
			return filtered
		}
	}
	return filtered
}
