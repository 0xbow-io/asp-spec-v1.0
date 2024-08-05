package feature

import (
	"github.com/swaggest/jsonschema-go"
)

type FeatureType struct {
	*jsonschema.Type
}

type Feature interface {
	T() FeatureType
	String() string
	Attributes() []interface{}
	Schema() *jsonschema.Schema
}

type FeatureAttribute interface {
	String() string
	TagType() string
	Tag(string) string
}
