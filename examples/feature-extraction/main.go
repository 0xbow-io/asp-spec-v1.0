package main

import (
	"fmt"

	// load the HIGH_RISK category schema
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/category/categories/HIGH_RISK"
)

func main() {
	if schema, err := HIGH_RISK_CATEGORY.MarshalJSON(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(schema))
	}
}
