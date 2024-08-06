# 6.3 Query Interface

**The `Record Category Registry` provides some public functions for querying the registry:**

```solidity

    /*//////////////////////////////////////////////////////////////////////////
    | PUBLIC FUNCTIONS
    //////////////////////////////////////////////////////////////////////////*/

    /**
     * @notice Returns the category bitmap for a record hash for a specific protocol scope
     * @param scope The protocol scope identifier
     * @param recordHash The hash of the record event
     * @return categoryBitmap The category bitmap for the record hash
     */
    function getCategoryBitmap(
        uint256 scope,
        bytes32 recordHash
    ) public view returns (bytes32 categoryBitmap) {
        (bool exists, bytes32 bitmap) = scopeRecordCategories[scope].tryGet(
            recordHash
        );
        if (!exists) {
            revert RecordNotFound(scope, recordHash);
        }
        return bitmap;
    }

    /**
     * @notice Returns the category bitmap for a record hash at a given index
     *          for a specific protocol scope
     * @param scope The protocol scope identifier
     * @param index The index of the record hash
     * @return recordHash recordHash at the given index
     * @return categoryBitmap The category bitmap for the record hash
     */
    function getRecordHashAndCategoryAt(
        uint256 scope,
        uint256 index
    ) public view returns (bytes32 recordHash, bytes32 categoryBitmap) {
        return scopeRecordCategories[scope].at(index);
    }

    /**
     * @notice Return the category bitmap for a record hash
     *          for a specific protocol scope
     # @dev does not revert if the record hash does not exist
     * @param scope The protocol scope identifier
     * @param recordHash The hash of the record event
     * @return exists A boolean indicating if the record hash exists
     * @return categoryBitmap The category bitmap for the record hash
     */
    function tryGetCategoryBitmap(
        uint256 scope,
        bytes32 recordHash
    ) public view returns (bool exists, bytes32 categoryBitmap) {
        return scopeRecordCategories[scope].tryGet(recordHash);
    }

    /**
     * @notice Returns the record hashes and their categories for a specific protocol scope between
     *         a given range
     * @param scope The protocol scope identifier
     * @param from The start index of the range
     * @param to The end index of the range
     * @return recordHashes The record hashes for the given range
     * @return categoryBitmaps The category bitmaps for the given range
     */
    function getRecordHashesAndCategories(
        uint256 scope,
        uint256 from,
        uint256 to
    )
        public
        view
        returns (
            bytes32[] memory recordHashes,
            bytes32[] memory categoryBitmaps
        )
    {
        require(
            from < to && to <= scopeRecordCategories[scope].length(),
            "Invalid range"
        );
        uint256 length = to - from;
        recordHashes = new bytes32[](length);
        categoryBitmaps = new bytes32[](length);

        for (uint256 i = 0; i < length; i++) {
            (recordHashes[i], categoryBitmaps[i]) = scopeRecordCategories[scope]
                .at(from + i);
        }
    }

    /**
     * @notice Returns the last record hash & it's category and the merkle-root
     *         for a specific protocol scope
     * @param scope The protocol scope identifier
     * @return root The merkle root for the protocol scope
     * @return recordHash The hash of the last known record event
     * @return categoryBitmap The category bitmap for the last known record event
     * @return index The index of the last known record event
     */
    function getLatestForScope(
        uint256 scope
    )
        public
        view
        returns (
            uint256 root,
            bytes32 recordHash,
            bytes32 categoryBitmap,
            uint256 index
        )
    {
        root = scopeRecordMerkleTrees[scope]._root();
        index = scopeRecordCategories[scope].length();
        if (index > 0) {
            (recordHash, categoryBitmap) = scopeRecordCategories[scope].at(
                index - 1
            );
        }
    }

```

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
