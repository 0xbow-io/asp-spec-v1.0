package main

import (
	"fmt"

	// load the HIGH_RISK category schema
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/category/categories/HIGH_RISK"
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature/extraction/extractors/HIGH_RISK"
)

func main() {
	// get the schema for this category
	if schema, err := HIGH_RISK_CATEGORY.MarshalJSON(); err != nil {
		fmt.Println(err)
	} else {
		// initialize the extractor to extract features from the schema
		var extractor = Init(schema)
	}
}
