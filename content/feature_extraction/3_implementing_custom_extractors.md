# 4.3 Implementing Feature Extractors

```admonish tip

Feature extraction can delegate / leverage specialized services via `plugins`, i.e:
- `Chainalysis` can extract `AML-Compliance` related features.
- An adpater for `Chainalaysis API` can be implemented as a `plugin`.
- Utility of this plugin allow feature extraction to leverage Chainalysis services.

Speialized feature-extraction logic / algorithms can be implemented as `gadgets`, i.e:

- `EVM-Call-Tracer` gadget for fast-tracing of internal calls.
- `Token-Tracker` gadget for tracking all token movements.



```

### 4.3.1 Considerations:

As Faults in feature extraction can have downstream impact in the classification process,
implementation of a feature extractor must take in account the
complexity of the feature extraction process to ensure that:

- The process is reliable and `deterministc` to avoid errors.
- THe process is `verifiable` to ensure that the extracted features are correct.
- The process is designed to be as fast as possible to avoid delays in the classification process.
- The process is designed to be as efficient as possible to avoid unnecessary resource consumption.
- The process is designed to be as secure as possible to avoid data leaks.
- The process is designed to be as scalable as possible to handle large volumes of data.

### 4.3.2 Inputs & Outputs:

To elimniate unwanted influence by other components, the feature-extractor is designed to:

- Accept only the Record as input.
- Be integrated with access-control systems with policies to restrict access
  to the feature-extraction configuration or runtime environment.
- Log any incoming data packets from external sources to ensure data integrity.

Outputs of the feature extraction process is serialized data that is encoded in a verifiable format.
The extractor should sign it's output to ensure that the data is not tampered with during transit.

The output should not be a stateful object such as a document or a database record consisting
of the extracted features. Instead it is a set of verfiable [JSON Patch](https://datatracker.ietf.org/doc/html/rfc6902) operations
that can be applied to a known & verified state to derive the extracted features, i.e.:

```json
[
  {
    "op": 1,
    "root": "0x010010",
    "path": "HIGH_RISK/DIRECT_SANCTIONED_ENTITY_EXPOSURE",
    "value": "0x0f001010",
    "merkle-proof": {}
  },
  {
    "op": 2,
    "root": "0x010010",
    "path": "HIGH_RISK/INDIRECT_SANCTIONED_ENTITY_EXPOSURE",
    "value": "0x0f001010",
    "merkle-proof": {}
  }
]
```

These operations is applied to a `default` feature-document which is a document containing the features for a category (per Schema)
but with default values (encrypted or encoded values).

This document is represented as a
[merkle-tree](https://docs.iden3.io/w3c/merklization/) where then merkle-proofs is used to verify data integrity.

### 4.3.3 Plugins & Gadgets:

The feature extraction process should be modularized to allow for easy extension, testing and maintenance.
Interfaces to external systems (i.e. Chaainalysis API) should be abstracted to allow dependency injection
and supporting testing of feature extraction process in isolation.

Below is an example of a feature-extractor implemented in go.
It utilises `plugins` for interfacing with external systems for feature extraction
and `gadgets` (i.e. interpreting EVM call-traces) for specialized feature extraction.

```go
package highRiskCat

import (
	"encoding/json"

	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature/extraction/storage"
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature/extraction/extractors"
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature/extraction/gadgets"
	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature/extraction/plugins"

	"github.com/swaggest/jsonschema-go"
)

var (
	// Plugin IDs
	plugins = []string{
		"PLUG_CA_01",
	}
	// Gadget IDs
	gadgets = []string{
		"GA_01",
	}
	// Cache IDs
	storages = []string{
		"DB_01",
	}
)

type feature struct {
	ID      string `json:"$id"`
	Minimum int    `json:"Minimum"`
	Maximum int    `json:"Maximum"`
	Type    string `json:"Type"`
	Default string `json:"Default"`
}

func applyFeatureSchema(feature *feature, spec *jsonschema.Schema) error {
	// quickest is to marshall then unmarshall
	b, error := spec.MarshalJSON()
	if error != nil {
		return error
	}
	return json.Unmarshal(b, feature)
}

type _Extractor struct {
	schema   []byte
	pluginCl PluginCl
	gadgetCl GadgetCl
	storageCl  StorageCl
}

var _ FeatureExtractor = (*_Extractor)(nil)

func Init(schema []byte) *_Extractor {
	ex := _Extractor{schema: schema}
	// init plugins
	for _, id := range plugins {
		if ex.pluginCl.Connect(id) != nil {
			return nil
		}
	}
	// init gadgets
	for _, id := range gadgets {
		if ex.gadgetCl.Connect(id) != nil {
			return nil
		}
	}
	// init cache
	for _, id := range storages {
		if ex.storageCl.Connect(id) != nil {
			return nil
		}
	}
	return &ex
}

// Implements FeatureExtractorMetadata interface
func (ex *_Extractor) Name() string             { return "HIGH_RISK_CATEGORY_EXTRACTOR" }
func (ex *_Extractor) Describe() string         { return "Extracting features for the HIGH_RISK Category" }
func (ex *_Extractor) Ver() string              { return "0.1.0" }
func (ex *_Extractor) Author() string           { return "0box.io" }
func (ex *_Extractor) License() string          { return "MIT" }
func (ex *_Extractor) URL() string              { return "github.com/0xbow.io/asp-v1.0/" }
func (ex *_Extractor) GetFeatureSchema() []byte { return ex.schema }

// Parses the Schmea to build a feature set
func (ex *_Extractor) featureSet() (set []feature) {
	var (
		category = jsonschema.Schema{}
	)
	if category.UnmarshalJSON(ex.schema) == nil {
		// extract features
		featureSet := category.Properties["features"].TypeObject.Items.SchemaArray
		set = make([]feature, len(featureSet))
		// iterate over features
		for i, feature := range featureSet {
			// apply feature schema
			applyFeatureSchema(&set[i], feature.TypeObject)
		}
	}
	return
}

func comparator(x [32]byte, y [32]byte) *Op { return nil }
func mtRoot(map[string][32]byte) [32]byte   { return [32]byte{} }
func (ex *_Extractor) sign(v []byte) []byte { return nil }
func (ex *_Extractor) ExtractFeatures(record Record) (out []byte) {
	return
}

```
