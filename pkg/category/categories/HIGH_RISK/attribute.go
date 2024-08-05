package highRiskCat

import (
	"fmt"
	"strings"

	. "github.com/0xbow-io/asp-spec-V1.0/pkg/feature"
)

type _FeatureAttribute uint

var _ FeatureAttribute = (*_FeatureAttribute)(nil)

const (
	invalid _FeatureAttribute = iota
	name
	required
	_type
	examples
	pattern
	_default
	maximum
	exclusiveMinimum
	minimum
	exclusiveMaximum
)

func (f _FeatureAttribute) String() string {
	return [...]string{
		"Invalid",
		"Name",
		"Required",
		"Type",
		"Examples",
		"Pattern",
		"Default",
		"Maximum",
		"exclusiveMinimum",
		"Minimum",
		"exclusiveMaximum"}[f]
}
func (f _FeatureAttribute) TagType() string {
	if f == invalid {
		return ""
	}
	var tagType string = "json"
	if f != name {
		tagType = strings.ToLower(f.String())
	}
	return tagType
}

func (f _FeatureAttribute) Tag(value string) string {
	if f == invalid {
		return ""
	}
	return fmt.Sprintf(`%s:"%s"`, f.TagType(), value)
}
