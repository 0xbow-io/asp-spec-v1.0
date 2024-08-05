# 6.2 Data Structure and Storage

The Public Registry uses two primary data structures for efficient storage and querying:

1. **Scope & Record Hash based Categories Storage**:

   - Implemented as:

     `mapping(uint256 scope => EnumerableMap.Bytes32ToBytes32Map RecordToCategory) scopeRecordCategories`

     - enables quick lookup of category bitmaps for a given record hash within a specific scope
     - `uint256 scope` is the scope identifier
     - `EnumerableMap.Bytes32ToBytes32Map RecordToCategory` is a mapping of record hashes to category bitmaps
     - `EnumerableMap.Bytes32ToBytes32Map` is imported from the
       [EnumerableMap openzeppelin library](https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/structs/EnumerableMap.sol)
     - `EnumerableMap.Bytes32ToBytes32Map` is utilised for easy iteration over the set of record hashes.

2. **Scope based Record Merkle Trees**:

   - Implemented as

     `mapping(uint256 scope => LeanIMTData recordTree) scopeRecordTrees`

     - Supports inclusion proof verification for a given record hash within a specific scope
     - `uint256 scope` is the scope identifier
     - `LeanIMTData recordTree` is the merkle-tree representation of the set of record hashes within the scope
     - `LeanIMTData` is imported from the
       [InternalLeanIMT zk-kit library](https://github.com/privacy-scaling-explorations/zk-kit.solidity/blob/main/packages/lean-imt/contracts/InternalLeanIMT.sol)
     - `InternalLeanIMT, LeanIMTData` provides a gas-optimized merkle-tree implementation.
