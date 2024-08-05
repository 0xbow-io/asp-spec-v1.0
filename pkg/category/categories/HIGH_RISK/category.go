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

var HIGH_RISK_CATEGORY = new(CategorySchema).Create(
	"HIGH_RISK",
	"Category for records associated to illicit activities",
	[]Feature{
		_DIRECT_SANCTIONED_ENTITY_EXPOSURE.Feature(),
		_INDIRECT_SANCTIONED_ENTITY_EXPOSURE.Feature(),
	})
