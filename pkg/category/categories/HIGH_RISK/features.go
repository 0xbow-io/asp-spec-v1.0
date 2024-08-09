package highRiskCat

import (
	"fmt"

	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature"
	"github.com/swaggest/jsonschema-go"
)

type _Feature uint

const (
	OFAC_LIST_MEMBERSHIP _Feature = iota
	FATF_LIST_MEMBERSHIP
	TRANSACTION_AMOUNT
)

var _ Feature = (*_Feature)(nil)

func (f _Feature) T() FeatureType {
	switch f {
	case OFAC_LIST_MEMBERSHIP:
		return FeatureType{
			Type: new(jsonschema.Type).WithSimpleTypes(jsonschema.Boolean),
		}
	case FATF_LIST_MEMBERSHIP:
		return FeatureType{
			Type: new(jsonschema.Type).WithSimpleTypes(jsonschema.Boolean),
		}
	case TRANSACTION_AMOUNT:
		return FeatureType{
			Type: new(jsonschema.Type).WithSimpleTypes(jsonschema.Number),
		}
	}
	return FeatureType{Type: new(jsonschema.Type)}
}

func (f _Feature) Feature() Feature {
	return &f
}

func (f _Feature) String() string {
	return [...]string{
		"OFAC_LIST_MEMBERSHIP",
		"FATF_LIST_MEMBERSHIP",
		"TRANSACTION_AMOUNT",
	}[f]
}

func (f _Feature) Attributes() []interface{} {
	return [...][]interface{}{
		{
			required: "true",
			_default: true,
		},
		{
			required: "true",
			_default: true,
		},
		{
			required: "true",
			_default: true,
		},
	}[f]
}

func (f _Feature) Schema(idPrefix string) (schema *jsonschema.Schema) {
	id := fmt.Sprintf("%s:features:%s", idPrefix, f.String())
	schema = &jsonschema.Schema{
		ID:   &id,
		Type: f.T().Type,
	}
	_attributes := f.Attributes()
	schema.WithDefault(_attributes[_default])
	return
}
