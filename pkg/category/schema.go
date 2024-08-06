package highRiskCat

import (
	"encoding/json"

	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature"
	"github.com/swaggest/jsonschema-go"
)

type CategorySchema struct {
	*jsonschema.Schema
}

func (s *CategorySchema) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(s.Schema, "", " ")
}

func (s *CategorySchema) applyFeatures(features []Feature) {
	for i, feature := range features {
		// add the feature schema to the array
		s.Schema.Properties["features"].TypeObject.Items.SchemaArray[i] =
			jsonschema.SchemaOrBool{TypeObject: feature.Schema()}
		s.Schema.Properties["features"].TypeObject.Required[i] = feature.String()
	}
}

func (s *CategorySchema) Create(name, description string, features []Feature) CategorySchema {
	s = &CategorySchema{
		Schema: &jsonschema.Schema{
			ID:          &name,
			Description: &description,
			Properties: map[string]jsonschema.SchemaOrBool{
				// category - feature schema
				"features": {
					TypeObject: &jsonschema.Schema{
						Required: make([]string, len(features)),
						Type:     new(jsonschema.Type).WithSimpleTypes(jsonschema.Array),
						// arary of features
						Items: &jsonschema.Items{
							SchemaArray: make([]jsonschema.SchemaOrBool, len(features)),
						},
					},
				},
			},
		},
	}
	defer s.applyFeatures(features)
	return *s
}
