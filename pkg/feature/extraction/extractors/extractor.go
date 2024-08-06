package extractor

/// The set of Features defined in the Feature Schema
/// Can be represented as a merkle-tree (Feature tree)
/// A patch operation describes Feature/s changes
/// And the resulting merkle-root of the
/// updated Feature Tree after applying the patch

type Op uint

const (
	OpAdd     Op = iota // Add a new value attribute in the Feature
	OpRemove            // Remove a value attribute in the Feature
	OpReplace           // Replace a value attribute in the Feature
	OpNull              // Nullify a value attribute in the Feature
)

type Patch struct {
	// The computed root of the Feature Set prior to the patch
	Root [32]byte
	// The operation applied
	Operation Op
	// Feature index
	FeatureIdx uint
	// The Feature
	Value [32]byte
}

// / FeatureExtractorMetadata
// / provides metadata about the Feature Extractor
type FeatureExtractorMetadata interface {
	Name() string
	Describe() string
	Ver() string
	Author() string
	License() string
	URL() string
	// getFeatureSchema returns
	// the serialized JSON schema for the Feature
	GetFeatureSchema() []byte
}

type Record interface {
	/// Record ID
	ID() []byte
	/// Raw Record Data
	Raw() []byte
}

type FeatureExtractor interface {
	FeatureExtractorMetadata
	// extractFeatures accepts a Record and
	// returns a set of signed patch operations
	ExtractFeatures(record Record) ([]Patch, [32]byte, [32]byte, error)
}
