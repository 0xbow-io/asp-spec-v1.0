# 4.3 Implementing Feature Extractors

```admonish tip
Feature Extractor implementation need to consider:
- Time taken to extract features
- Computation cost of extracting features
- The complexity of verifying the features

It is recommended that the Extractor is designed for
singular extraction and delegate extraction to specialized services
when necessary.
```

When implementing extractors, ensure that:

1. Feature extraction are deterministic for a given record.
2. Feature extraction is computationally efficient.
3. Designed with a modular approach to allow for easy integration with other services.

Here is an example go-implementation of feature-extractor that utilises a Chainalaysis plugin for
extracting featues relevant to the `HIGH_RISK` category.

```go

package highRisk

import (
	feature "github.com/0xbow.io/asp-v1.0/category-engine/feature"
	category "github.com/0xbow.io/asp-v1.0/category-engine/category"
	extractor "github.com/0xbow.io/asp-v1.0/category-engine/feature/extraction/extractor"
	plugin "github.com/0xbow.io/asp-v1.0/category-engine/feature/extraction/plugin"
	Record "github.com/0xbow.io/asp-v1.0/observer/types/record"
)

var categoryFeatureSchema string = `{
  "title": "HIGH_RISK",
  "description": "HIGH_RISK is a category for records associated to illicit activities ",
  "required": [
    "DIRECT_SANCTIONED_ENTITY_EXPOSURE",
    "INDIRECT_SANCTIONED_ENTITY_EXPOSURE"
  ],
  "properties": {
    "DIRECT_SANCTIONED_ENTITY_EXPOSURE": {
      "default": "0.0",
      "examples": [1000.1],
      "minimum": 1000000.01,
      "maximum": 0,
      "pattern": "0",
      "type": "number"
    },
    "INDIRECT_SANCTIONED_ENTITY_EXPOSURE": {
      "default": "0.0",
      "examples": [1000.1],
      "minimum ": 1000000.01,
      "maximum": 0,
      "pattern": "0",
      "type": "number"
    }
  },
  "type": "object"
}`


type _Extractor struct {
	schena category.Schema
	pluginCl plugin.Client   `feature:"plugin"`
	conf     *feature.Config `feature:"config"`
	logger   feature.Logger  `feature:"logger"`
}
var _ extractor.FeatureExtractor = (*_Extractor)(nil)


func (*_Extractor) Init(conf *feature.Config){
	var err error
	ex := _Extractor{
		conf:   conf,
		logger: conf.Logger,
	}

	ex.pluginCl, err = plugin.NewClient(conf, ex.getPluginList())
	if err != nil {
		return nil, err
	}
	return &ex, nil
}

func (ex _Extractor) getName() string        { return "HIGH_RISK_FEATURES_EXTRACTOR" }
func (ex _Extractor) getDescription() string { return "Extracting features for the HIGH_RISK Category" }
func (ex _Extractor) getVersion() string     { return "0.1.0" }
func (ex _Extractor) getAuthor() string      { return "0box.io" }
func (ex _Extractor) getLicense() string     { return "MIT" }
func (ex _Extractor) getURL() string         { return "github.com/0xbow.io/asp-v1.0/" }
func (ex _Extractor) getSupportedFeatures() []feature.FeatureType {
	return []feature.FeatureType{feature.Categorical}
}
func (ex _Extractor) getPluginList() []string { return []string{"chainalaysis"} }

func (ex _Extractor) extractFeatures(record Record.Record) ([]feature.Patch, sig [32]byte error) {
	// Compose request to plugin
	// Parse & validate response
	// Return features
	return nil, nil
}

```
