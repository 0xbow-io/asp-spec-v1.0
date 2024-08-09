# 5.1 Rule-Based Classification

```admonish info title='Who set the rules?'

The responsibility of defining a category may be assigned to either an individual entity
or a collaboration of entities such as:

- The ASP enttity
- The protocol entity
- The network governance entity
- The regulatory entity


```

## 5.1.1 Overview

The classification process evaluates records with a set of rules to determine it's categories.

The ASP considers that an object `categorized` as $\varGamma_n$ is therefore $\varGamma_n$ `compliant `and must have `satisfied`
all the `rules` of $\varGamma_n$, i.e.:

- A transaction is categorised as `AML` compliant because it has satisfied `AML rules` such as `Sender is not in the OFAC list`.
- A person is categorised as `KYC` compliant because it has satisfied `KYC rules` such as `Person has provided a valid ID`.
- A vote is categorised as `Valid Vote` because it has satisfied `voting rules` such as `Vote is casted within the voting period`.
- A document is categorised as `Approved Document` because it has satisfied
  `document rules` such as `Document is signed & reviewed by an auditor`.

### 1. The Compliance Domain:

Given a `compliance` predicate $P$ and it's propositional function:
$$P(x,\Gamma) = x \text{ is } \Gamma \text{ compliant }$$

- Where $\Gamma$ is a set of atomic rules:
  $$\Gamma = \{r_1, r_2, r_3, \ldots, r_n\}$$

$\Gamma$ is considered a `Compliance Domain` only if the following holds true:

$$P(x,\Gamma) \iff \forall r \in \Gamma : r(x) = true$$
$$P(y,\Gamma) \iff \exists r \in \Gamma : r(y) = false$$
$$ x \in X, y \in Y, X \subset Z, Y \subset Z$$

Where:

- $x$ is compliant with $\Gamma$ if and only if $x$ satisfies all the rules in $\Gamma$.
- $y$ is not compliant with $\Gamma$ if and only if $y$ does not satisfy all the rules in $\Gamma$.
- $X$ is a set of transactions that are compliant with $\Gamma$.
- $Y$ is a set of transactions that are not compliant with $\Gamma$.
- $X$ and $Y$ are subsets of the universal set $Z$.

Example:

If the statements are true:

- A set of transactions ($X$) exists such that each transaction ($x$) is AML ($\Gamma$) compliant.
- A set of transactions ($Y$) exists such that each transaction ($y$) is not AML ($\Gamma$) compliant.
- $X$ and $Y$ belong to a super set of transactions $Z$.

Then `Compliance Domain` ($\Gamma$) exists as a set of rules for
which holds true for all transactions in $X$ and false for all transactions in $Y$.

### 2. Defining Compliance Rules:

By identifiying ($\Gamma$) and set $X$ & $Y$
it is then possible to derive the compliance rules which governs $P$:

$$\Gamma = \{r_1, r_2, r_3, \ldots, r_n\}$$

- I.e., The rules for `AML` compliance could be:

  - Sender & receiver is not in the OFAC list
  - Sender & receiver is not in the FATF list
  - Transaction amount is less than 1000 USD.

- These rules are verified against a set of:
  - `AML` compliant transactions ($X$)
  - `non-AML` compliant transactions ($Y$).

```admonish tip Title='Keep it simple'

Rules should be represented as `atomic` & `independent` and therefore avoid any:

- Representations indicating hierarchical structure of rulesi.e.:
  - $R:$ Transaction is not from a sanctioned entity
    - $r_i:$ Sender is not in the OFAC list

- Representations indicating dependencies between rules i.e.:
  - $R:$ Transaction is not from a sanctioned entity
    - $r_i:$ Sender is not in the OFAC list
    - $r_j:$ Sender is not in the FATF list


```

### 2. Deriving Categories:

A category is a translation of $P(x,\Gamma)$ where`features` & `thresholds`
represents the atomic rules $r_i$ of $\Gamma$:

- $\Gamma:$ `AML`
  - Category: `AML_COMPLIANT`
- $r_1:$ Sender & receiver is not in the OFAC list
  - Feature: `OFAC_LIST_MEMBERSHIP`
  - Threshold: `false`
  - `OFAC_LIST_MEMBERSHIP` is `false` for a record.
- $r_1:$ Sender & receiver is not in the FATF list
  - Feature: `FATF_LIST_MEMBERSHIP`
  - Threshold: `false`
  - `FATF_LIST_MEMBERSHIP` is `false` for a record.
- $r_1:$ Transaction amount is less than 1000 USD.
  - Feature: `TRANSACTION_AMOUNT`
  - Threshold: `1000`
  - `TRANSACTION_AMOUNT` is less than `1000` for a record.

This translation (see [4.2 Feature Types and Formats](/feature_extraction/1_feature_extractor_interface.md)) results in a
`JSON Schema` document such as:

```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "tag:0xbow.io,2024:categories:AML_COMPLIANT",
  "title": "Record is AML Compliant",
  "type": "object",
  "properties": {
    "features": {
      "type": "object",
      "properties": {
        "OFAC_LIST_MEMBERSHIP": {
          "$id": "tag:0xbow.io,2024:categories:AML_COMPLIANT:features:OFAC_LIST_MEMBERSHIP",
          "type": "boolean",
          "default": "true"
        },
        "FATF_LIST_MEMBERSHIP": {
          "$id": "tag:0xbow.io,2024:categories:AML_COMPLIANT:features:FATF_LIST_MEMBERSHIP",
          "type": "boolean",
          "default": "true"
        },
        "TRANSACTION_AMOUNT": {
          "$id": "tag:0xbow.io,2024:categories:AML_COMPLIANT:features:TRANSACTION_AMOUNT",
          "type": "integer",
          "default": "1000"
        }
      },
      "required": [
        "OFAC_LIST_MEMBERSHIP",
        "FATF_LIST_MEMBERSHIP",
        "TRANSACTION_AMOUNT"
      ]
    }
  },
  "required": ["features"]
}
```

```admonish note Title='Feature Types'

Features can be considered as the properties or attributes of a record which is later
evaluated against the compliance rules.

i.e. The rule `Sender & receiver is not in the OFAC list` evaluates
the value of 'OFAC_LIST_MEMBERSHIP' which is of a `boolean` type.

There are no constraints on the type of features that can be used in the schema.
It could be a `string`, `integer`, `boolean`, `array`, `object` etc.

However it should not include or point/reer to another feature
nor be a function of another feature.

```

### 3. Classifying a record :

**Note that default values are set to invalidate the category for the record.**

The `default` document is a document which satisfy the Schema but where all `features` have values
set to its `default` value.

```json
{
  "features": {
    "OFAC_LIST_MEMBERSHIP": true,
    "FATF_LIST_MEMBERSHIP": true,
    "TRANSACTION_AMOUNT": 1000
  }
}
```

As outlined in [section 4.3](/feature_extraction/3_implementing_custom_extractors.md),
the `Feature Extractor` delivers the values of these features in the form of `JSON patch` operations.

```json
"patch": [
  {
    "op": 2,
    "root": "0x010010",
    "$id": "tag:0xbow.io,2024:categories:AML_COMPLIANT:features:OFAC_LIST_MEMBERSHIP",
    "value": "false",
    "merkle-proof": {}
  },
  {
    "op": 2,
    "root": "0x010010",
    "$d": "tag:0xbow.io,2024:categories:AML_COMPLIANT:features:FATF_LIST_MEMBERSHIP",
    "value": "false",
    "merkle-proof": {}
  },
  {
    "op": 2,
    "root": "0x010010",
    "$d": "tag:0xbow.io,2024:categories:AML_COMPLIANT:features:TRANSACTION_AMOUNT",
    "value": "100",
    "merkle-proof": {}
  }
]
```

If when applying these patches to the `default` document, the document still satisfies the schema
and the `features` are consistent with the rules,
then the record is classified as `AML_COMPLIANT`.

```json
{
  "features": {
    "OFAC_LIST_MEMBERSHIP": false,
    "FATF_LIST_MEMBERSHIP": false,
    "TRANSACTION_AMOUNT": 100
  }
}
```
