package highRiskCat

import (
	"encoding/json"

	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature/extraction/caches"
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
	caches = []string{
		"CACHE_REDIS",
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
	cacheCl  CacheCl
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
	for _, id := range caches {
		if ex.cacheCl.Connect(id) != nil {
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
func (ex *_Extractor) ExtractFeatures(record Record) (patches []Patch, root [32]byte, sig [32]byte, err error) {
	return
}
