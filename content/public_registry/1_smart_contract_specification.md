# 6.1 Smart Contract Specification

```admonish info
The Public Registry is a collection of smart contracts which serves as the on-chain storage solution for the ASP.
It provides the necessary interfaces for onchain protocols to integrate with the ASP.
```

The current [registry](https://github.com/0xbow-io/asp-contracts-V1.0) is composed of 2 core contracts:

- **Record Category Registry** ([view contract](https://github.com/0xbow-io/asp-contracts-V1.0/blob/main/src/RecordCategoryRegistry.sol))
- **Association Set Provider** ([view contract](https://github.com/0xbow-io/asp-contracts-V1.0/blob/main/src/RecordCategoryRegistry.sol))

```mermaid

---
title: "Figure 6.1: ASP Smart Contract Class Diagram"
---

%%{
  init: {
    'theme': 'base',
    'themeVariables': {
      'primaryColor': '#1e1e2e',
      'primaryTextColor': '#cdd6f4',
      'primaryBorderColor': '#89b4fa',
      'lineColor': '#fab387',
      'secondaryColor': '#181825',
      'tertiaryColor': '#1e1e2e',
      "clusterBorder": "#f2cdcd",
      'noteTextColor': '#f5e0dc',
      'noteBkgColor': '#f5c2e7',
      'notesBorderColor': '#cba6f7',
      'textColor': '#f5e0dc',
      'fontSize': '16px',
      'labelTextColor': '#f5e0dc',
      'actorBorder': '#89b4fa',
      'actorBkg': '#1e1e2e',
      'actorTextColor': '#f5e0dc',
      'actorLineColor': '#89b4fa',
      'signalColor': '#cdd6f4',
      'signalTextColor': '#f5e0dc',
      'messageTextColor': '#f5e0dc',
      'messageLine0TextColor': '#f5e0dc',
      'messageLine1TextColor': '#f5e0dc',
      'loopTextColor': '#f5e0dc',
      'activationBorderColor': '#f5c2e7',
      'activationBkgColor': '#1e1e2e',
      'sequenceNumberColor': '#1e1e2e'
    }
  }
}%%


classDiagram

  class AccessControl{
    ~Map~bytes32|Set~address~~ _roleMembers
    +grant(account)
    +revoke(role,account)
    +has(role,account)
  }


  class Registry {
    +bytes32 REGISTRY_ADMIN_ROLE
    -Map~uint256|MerkleTree~bytes32~~ _scope_record_trees
    -Map~uint256|Map~bytes32|bytes32~~ _scope_record_categories

    +setRecordCategory(scope,r,c)
    +getCategoryForRecord(scope,r)
    +getRecordAndCategoryAt(scope,index)
    +getRecordsAndCategories(scope, from, to)
    +getLatestForScope(scope)

  }

  class ASP{
    +applyPredicate(scope, records, categoryMask, predicateType)
    -_applyPredicate(predicateType, categoryMask, categoryBitmap)
  }

  class PredicateType{
      <<enumeration>>
      Intersection
      Union
      Complement
  }


  AccessControl --|> Registry : Inheritance
  Registry --|> ASP : Inheritance

  ASP "1" -- "*" PredicateType: contains
```
