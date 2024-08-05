# 6.3 Query Interface

The query interface allows external systems to retrieve subsets of record hashes

based on scope, category, and predicate type. The `queryRecords` function implements this interface:

```solidity

function _applyPredicate(
    PredicateType predicateType,
    bytes32 characteristicFunction,
    bytes32 elementProperties
) internal pure returns (bool satisfiesPredicate) {
    if (predicateType == PredicateType.Intersection) {
        satisfiesPredicate =
            (elementProperties & characteristicFunction) ==
            characteristicFunction;
    } else if (predicateType == PredicateType.Union) {
        satisfiesPredicate =
            (elementProperties & characteristicFunction) != 0;
    } else if (predicateType == PredicateType.Complement) {
        satisfiesPredicate =
            (elementProperties & characteristicFunction) == 0;
    }
}

function applyPredicate(
    uint256 domain,
    bytes32[] calldata subset,
    bytes32 characteristicFunction,
    PredicateType predicateType
) public view returns (bytes32[] memory elements, uint256 setCardinality) {
    bytes32[] memory satisfyingElements = new bytes32[](subset.length);
    for (uint256 i = 0; i < subset.length; i++) {
        bytes32 element = subset[i];
        (bool isMember, bytes32 elementProperties) = tryGetCategoryBitmap(
            domain,
            element
        );
        if (!isMember) {
            continue;
        }
        if (
            _applyPredicate(
                predicateType,
                characteristicFunction,
                elementProperties
            )
        ) {
            satisfyingElements[setCardinality] = element;
            setCardinality++;
        }
    }
    assembly {
        mstore(satisfyingElements, setCardinality)
    }
    return (satisfyingElements, setCardinality);
}

```

Predicate types:

- 0: Intersection (all bits in `categoryMask` must be set)
- 1: Union (at least one bit in `categoryMask` must be set)
- 2: Complement (no bits in `categoryMask` should be set)
