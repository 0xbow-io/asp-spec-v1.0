# 4.1 Feature Extractor Interface

> Feature-extraction refers to the process of transforming raw-data into numerical features that can
> be processed whilst preserving the information in the original data set.

Feature Extraction is `category` driven whereby features are `properties`, `attributes`,
or `characteristics` of a Record that are meaningfuly grouped together to form a `category`:

- **Examples Features:**

  - Accounts associated with the Record beloging to a certain list, i.e. wallet address that executed the transaction:
    - Feature: `ACCOUNT_BLACKLISTED`
  - Asset exposure to certain sources or activities such as Gambling, Money Laundering, etc.
    - Feature: `ASSET_EXPOSURE_MONYLAUNDERING`
  - Volume of Assets transferred by a particular list of accounts
    - Feature: `ACCOUNTLIST_ASSET_TRANSFER_VOLUME`
  - Time of the day when the transaction was executed
    - Feature: `TRANSACTION_TIME`
  - If a certain internal call went against a specific Policy
    - Feature: `POLICY_0X1A0B_VIOLATION`

- **Example Categories:**

  - `UNAPPROVED` category that groups the following features:
    - `POLICY_0X1A0B_VIOLATION`
      - **Threshold**: True
    - `ACCOUNTLIST_ASSET_TRANSFER_VOLUME`
      - **Threshold**: 1000
  - `APPROVED` category that groups the following features:
    - `ACCOUNT_BLACKLISTED`
      - **Threshold**: False
    - `ASSET_EXPOSURE_MONYLAUNDERING`
      - **Threshold**: False
  - `SUSPICIOUS` category that groups the following features:
    - `ACCOUNT_BLACKLISTED`
      - **Threshold**: True
    - `ASSET_EXPOSURE_MONYLAUNDERING`
      - **Threshold**: True

The Feature Extractor is responsible for extracting these features from the Record and delivering
them in a structured format. It's interface should be minimalistic and easy to implement.

Below is an example interface for the Feature Extractor:

```go
package extractor

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
	// extractFeatures extracts
	// features from the given Record
	extractFeatures([]byte) ([]byte)
}

```

The `extractFeatures` function is the core method which processes a given `Record` object and returns data
required for the classification process.
