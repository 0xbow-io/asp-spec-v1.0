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

The ASP follows a rule-based approach in defining categories & it's features and can be summarised as follows:

### 1. Discovering the Compliance Domain Space:

The `Compliance Domain Space` $X$ is defined as a set of $x$ where all values of $x$ satisfies compliance Predicate $P(x)$:

$$\forall x \in X , P(x)$$

For example: $\forall x \in X , P(x)$ = transaction ($x$) is AML compliant.

Discovering $X$ would simply require creating a set of transactions that are AML compliant.

### 2. Deriving Compliance Rules:

Once $X$ is discovered, it is possible to derive the conditional rules that governs $P(x)$ by
identifying the common properties between all values of $x \in X$.

For example: $\forall x \in X , P(x)$ = transaction ($x$) is AML compliant.

The common properties could be:

- Transaction amount is less the acceptable threshold defined by the regulatory entity.
  - Threshold amount is $1000 for AML compliance in the US.
- Transaction is not from a high-risk country
  - High-risk country is defined as a country that has been flagged by the regulatory entity
    - i.e. Receiver & Sender are not in a FATF list.
- Transaction is not from a high-risk entity
  - High-risk entity is defined as an entity that has been flagged by the regulatory entity
    - i.e. Receiver & Sender are not in a OFAC list.

Applying a reductionist approach to the common properties, the compliance rules can be simpliefied to:

- Receiver & Sender are not members of the OFAC list
- Receiver & Sender are not members of the FATF list
- Transaction amount is less than $1000

### 2. Deriving Categories:

The Category is a translation of the compliance predicate $P(x)$, i.e. "AML compliant" and its' features are
the compliance rules that governs $P(x)$. As shown in section
[4.2 Feature Types and Formats](/feature_extraction/1_feature_extractor_interface.md), this translation is
conveyed in the form of a `JSON Schema` document such as:

```json
{
  "$id": "AML_COMPLIANT",
  "description": "Category for AML compliance",
  "properties": {
    "features": {
      "items": [
        {
          "$id": "OFAC_LIST_MEMBERSHIP",
          "default": true,
          "examples": ["false"],
          "maximum": true,
          "minimum": false,
          "pattern": "",
          "type": "boolean"
        },
        {
          "$id": "FATF_LIST_MEMBERSHIP",
          "default": true,
          "examples": ["false"],
          "maximum": true,
          "minimum": false,
          "pattern": "",
          "type": "number"
        },
        {
          "$id": "TRANSACTION_AMOUNT",
          "default": 3000000,
          "examples": ["1000"],
          "maximum": 1000,
          "minimum": 0,
          "pattern": "",
          "type": "number"
        }
      ],
      "required": ["OFAC_LIST_MEMBERSHIP", "FATF_LIST_MEMBERSHIP"],
      "type": "array"
    }
  }
}
```

**Note in the schema the default values are set to not satisfy the thresholds.**
For example, the `default` feature document for $x$:

```json
{
  "$id": "AML_COMPLIANT",
  "features": [
    {
      "OFAC_LIST_MEMBERSHIP": true
    },
    {
      "FATF_LIST_MEMBERSHIP": true
    },
    {
      "TRANSACTION_AMOUNT": 3000000
    }
  ]
}
```

$x$ will only satisfy $P(x)$ if given the correct set of [Patch](/feature_extraction/3_implementing_custom_extractors.md)
operations which would set the values of the features to satisfy the compliance rules:

```json
"patch": [
  {
    "op": 1,
    "root": "0x010010",
    "path": "AML_COMPLIANT/OFAC_LIST_MEMBERSHIP",
    "value": "false",
    "merkle-proof": {}
  },
  {
    "op": 2,
    "root": "0x010010",
    "path": "AML_COMPLIANT/FATF_LIST_MEMBERSHIP",
    "value": "false",
    "merkle-proof": {}
  }.
  {
    "op": 2,
    "root": "0x010010",
    "path": "AML_COMPLIANT/TRANSACTION_AMOUNT",
    "value": "100",
    "merkle-proof": {}
  }
]
```
