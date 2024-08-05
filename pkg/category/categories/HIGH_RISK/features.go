package highRiskCat

import (
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature"
	"github.com/swaggest/jsonschema-go"
)

type _Feature uint

const (
	_DIRECT_SANCTIONED_ENTITY_EXPOSURE _Feature = iota
	_INDIRECT_SANCTIONED_ENTITY_EXPOSURE
)

var _ Feature = (*_Feature)(nil)

func (f _Feature) T() FeatureType {
	return FeatureType{Type: new(jsonschema.Type).WithSimpleTypes(jsonschema.Object)}
}

func (f _Feature) Feature() Feature {
	return &f
}

func (f _Feature) String() string {
	return [...]string{"DIRECT_SANCTIONED_ENTITY_EXPOSURE", "INDIRECT_SANCTIONED_ENTITY_EXPOSURE"}[f]
}

func (f _Feature) Attributes() []interface{} {
	return [...][]interface{}{
		{
			name:             f.String(),
			required:         "true",
			examples:         "1000.1",
			pattern:          "",
			_default:         0.0,
			maximum:          0.0,
			exclusiveMinimum: 0.0,
			minimum:          1000000.01,
			exclusiveMaximum: 0.0,
		},
		{
			name:             f.String(),
			required:         "true",
			examples:         "1000.1",
			pattern:          "",
			_default:         0.0,
			maximum:          0.0,
			exclusiveMinimum: 0.0,
			minimum:          1000000.01,
			exclusiveMaximum: 0.0,
		},
	}[f]
}

func (f _Feature) Schema() (schema *jsonschema.Schema) {
	id := f.String()
	schema = &jsonschema.Schema{
		ID:   &id,
		Type: f.T().Type,
	}

	_attributes := f.Attributes()
	schema.WithDefault(_attributes[_default])
	schema.WithExamples(_attributes[examples])
	schema.WithPattern(_attributes[pattern].(string))
	schema.WithMaximum(_attributes[maximum].(float64))
	schema.WithExclusiveMaximum(_attributes[exclusiveMaximum].(float64))
	schema.WithMinimum(_attributes[minimum].(float64))

	return
}
