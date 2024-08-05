# 4.2 Feature Types and Formats

Feature Extraction is `category` driven which means that features are `properties`, `attributes`,
or `characteristics` of a Record that are only relevant to the category.

For example, the `KYC_PASS` category for records that are associated with an account that has passed KYC might depend
on whether the an `KYC attestation` (i.e. Coinbase EAS attestation) for the `associated account` exists.

Therefore we can define the category feature-set as:

- Feature 1: `ASSOCIATED_ACCOUNT` - The address of the associated account.
- Feature 2: `KYC_ATTESTATION` - The KYC attestation status of the account.

The classification for `KYC_PASS` depends on the feature-extractor's ability to extract these features.

## 4.2.1 The Category Feature Schema

Features should be represented in a standardized & structured and verfiable format to ensure compatibility
across different components of the ASP system & facilitate interoperability with external systems.

Category Feature Schema is a [JSON Schema](https://www.learnjsonschema.com/2020-12/) document
that defines features attributes for all features belonging to a category, i.e:

```json
{
  "$id": "HIGH_RISK",
  "description": "Category for records associated to illicit activities",
  "properties": {
    "features": {
      "items": [
        {
          "$id": "DIRECT_SANCTIONED_ENTITY_EXPOSURE",
          "default": 0,
          "examples": ["1000000.1"],
          "maximum": 0,
          "exclusiveMaximum": 0,
          "minimum": 1000000.01,
          "pattern": "",
          "type": "object"
        },
        {
          "$id": "INDIRECT_SANCTIONED_ENTITY_EXPOSURE",
          "default": 0,
          "examples": ["3000000.1"],
          "maximum": 0,
          "exclusiveMaximum": 0,
          "minimum": 3000000.01,
          "pattern": "",
          "type": "object"
        }
      ],
      "required": [
        "DIRECT_SANCTIONED_ENTITY_EXPOSURE",
        "INDIRECT_SANCTIONED_ENTITY_EXPOSURE"
      ],
      "type": "array"
    }
  }
}
```

In the above example, the `HIGH_RISK` category has two features:

- feature 1: `DIRECT_SANCTIONED_ENTITY_EXPOSURE`

  - the attributes of this feature:

    - `default`: 0.0
    - `examples`: ["1000000.1"],
    - `minimum`: 1000000.01,
    - `maximum`: 0
    - `pattern`: ""
    - `type`: "number"

    - Based on the attributes, the value of the feature is of the `number` type and must be >1000000.01
      and a default value of 0.0

- feature 2: `INDIRECT_SANCTIONED_ENTITY_EXPOSURE`

  - the attributes of this feature:

    - `default`: 0.0
    - `examples`: ["3000000.1"],
    - `minimum`: 3000000.01,
    - `maximum`: 0
    - `pattern`: ""
    - `type`: "number"

  - Based on the attributes, the value of the feature is of the `number` type and must be >3000000.01
    and a default value of 0.0

JSON Schema's like the one above can be created by a schema generator.
Below is an example walkthrough in creating a schmea generater in go using the
[swaggest/jsonschema-go]("github.com/swaggest/jsonschema-go") package:

- ### Declaring the interfaces for Features:

```go
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


```

- ### Declaring Feature Schema for a Category:

```go
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
	return [...]string{
		"DIRECT_SANCTIONED_ENTITY_EXPOSURE",
		"INDIRECT_SANCTIONED_ENTITY_EXPOSURE"}[f]
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

```

- ### Category Schema Generator:

```go
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
		s.Schema.Properties["features"].TypeObject.Items.SchemaArray[i] = jsonschema.SchemaOrBool{TypeObject: feature.Schema()}
		s.Schema.Properties["features"].TypeObject.Required = append(s.Schema.Properties["features"].TypeObject.Required, feature.String())
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
						Type: new(jsonschema.Type).WithSimpleTypes(jsonschema.Array),
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
```

- ### Defining the Category - Feature schema for HIGH_RISK category:

```go
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

```
