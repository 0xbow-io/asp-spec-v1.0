# 4.1 Feature Extractor Interface

The Feature Extractor is a crucial component of the ASP system,
responsible for deriving meaningful attributes from state transition records that are
relevant for categories.

Example interface for the Feature Extractor:

```go
package extractor

/// The set of Features defined in the Feature Schema
/// Can be represented as a merkle-tree (Feature tree)
/// A patch operation describes Feature/s changes
/// And the resulting merkle-root of the
/// updated Feature Tree after applying the patch

type Op uint
const (
	OpAdd Op = iota
	OpRemove
	OpReplace
)

type Patch struct {
	// The computed root of the Feature tree prior to the patch
	Root [32]byte
	// The operation to be applied
	Operation Op
	// The index of the Feature to be updated, removed, added
	FeatureIdx uint
	// The new value of the Feature
	Value [32]byte
}


/// FeatureExtractorMetadata
/// provides metadata about the Feature Extractor
type FeatureExtractorMetadata interface {
	getName() string
	getDescription() string
	getVersion() string
	getAuthor() string
	getLicense() string
	getURL() string
	// getFeatureSchema returns
	// the serialized JSON schema for the Feature
	getFeatureSchema() ([]byte)
	getPluginList() ([]string)
}


type FeatureExtractor interface {
	FeatureExtractorMetadata
	// extractFeatures accepts a Record and
	// returns a set of patch operations
	extractFeatures(record Record) ([]Patch, sig [32]byte, error)
}

```

```admonish important
Feature extraction can be delegated to speacialized services via plugins / adapters, i.e:
- Chainalysis is a feature extractor for AML compliance related features.
- An adpater to Chainalaysis API can be implemented as a extraction plugin.
- Utility of this plugin allow feature extraction to leverage Chainalysis service.
```

The `extractFeatures` function is the core method that accepts a `Record` and returns a set of `Patch` operations,
along with a signature of the set of `Patch` operations:
  - The `Patch` operations describe the changes to be made to the Record Features.
  - Features that the Feature Extractor is capable of extracting are defined in the Feature Schema.
