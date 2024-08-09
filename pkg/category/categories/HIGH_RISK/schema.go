package highRiskCat

import (
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/category"
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature"
	"github.com/swaggest/jsonschema-go"
)

type _CategorySchema interface {
	MarshalJSON() ([]byte, error)
	Create(title, description string, features []jsonschema.Field) *CategorySchema
}

var AML_COMPLIANT = new(CategorySchema).Create(
	// Label
	"AML_COMPLIANT",
	// title
	"Record is AML Compliant",
	// Required Features
	[]Feature{
		Feature(OFAC_LIST_MEMBERSHIP),
		Feature(FATF_LIST_MEMBERSHIP),
		Feature(TRANSACTION_AMOUNT),
	})
