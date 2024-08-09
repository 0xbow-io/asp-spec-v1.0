package highRiskCat

import (
	"encoding/json"
	"fmt"

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
		s.Schema.Properties["features"].TypeObject.Properties[feature.String()] = jsonschema.SchemaOrBool{
			TypeObject: feature.Schema(*s.Schema.ID),
		}
		s.Schema.Properties["features"].TypeObject.Required[i] = feature.String()
	}
}

func (s *CategorySchema) Generate(
	label,
	title string,
	features []Feature) CategorySchema {
	id := fmt.Sprintf("0xbow.io,2024:categories:%s", label)
	s = &CategorySchema{
		Schema: &jsonschema.Schema{
			ID:    &id,
			Title: &title,
			Properties: map[string]jsonschema.SchemaOrBool{
				// category - feature schema
				"features": {
					TypeObject: &jsonschema.Schema{
						Required:   make([]string, len(features)),
						Type:       new(jsonschema.Type).WithSimpleTypes(jsonschema.Object),
						Properties: make(map[string]jsonschema.SchemaOrBool),
					},
				},
			},
		},
	}
	defer s.applyFeatures(features)
	return *s
}
