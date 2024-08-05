# 5.1 Rule-Based Classification

```admonish info title= "Who sets the rules?"
Classification rules can be set by:

1. The ASP Entity (i.e. 0xBow.io)
2. Individual protocols integrated with the ASP
3. Governance of the network for which the ASP operates in.
4. Regulatory bodies active in the region of ASP operations.
4. Or A collaboration of protocols, network governance, regualtors and ASP.
```

In ASP V1.0, Classifiers apply a series of conditional statement to map extracted features to specific categories.
Example of a rule-based classifier:

```go

func (cl *AMLClassifier) classify(features []feature.Feature) []Category {
	var (
		cats = make([]Category, len(features))
	)

	for i, f := range features {
		switch f.FeatureType {
		case feature.BOOLEAN:
			if f.Name == "kyc" {
				if !f.Value.(*boolean) {
					cats[i] = MISSING_KYC
				} else {
					cats[i] = KNOWN_KYC
				}
			} else {
				cats[i] = UNKNOWN_NAME
			}
		case feature.CATEGORICAL:
			if f.Name == "illicit_exposure" {
				if f.Value.(*string) == "high" {
					cats[i] = HIGH_RISK
				} else if f.Value.(*string) == "medium" {
					cats[i] = MEDIUM_RISK
				} else if f.Value.(*string) == "low" {
					cats[i] = LOW_RISK
				} else {
					cats[i] = UNKNOWN_RISK
				}
			} else {
				cats[i] = UNKNOWN_EXPOSURE
			}
		default:
			cats[i] = UNKNOWN_FEATURE
		}
	}
	return cats
}
```
