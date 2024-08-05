# 5.4 Categorization Best Practices

```mermaid

---
title: "Figure 5.1: Heirarchical Categorization"
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

graph TD
    AML["AML (topic)"] --> Global["Global (scope)"]
    Global  --> KYC["KYC (category)"]
    KYC  --> Present["Present (Threshold)"]
     KYC  --> Missing["Missing (Threshold)"]
    AML  --> Exposure["Exposure (scope)"]
    Exposure --> Illicit["Illicit (category)"]
    Illicit --> High["High (Threshold)"]
    Illicit --> Medium["Medium (Threshold)"]
    Illicit --> Low["Low (Threshold)"]
```

1. **Hierarchical Categories**: Organize categories in a hierarchical structure for more nuanced classification.

2. **Multi-Label Classification**: Allow records to belong to multiple categories simultaneously.

3. **Threshold Scores**: Include threshold scores for each category assignment.

4. **Interpretability**: Explain category assignments.

5. **Cross-Protocol Consistency**: Ensure consistent categorization across different protocols for similar state transitions.

6. **Version Control**: Maintain version control for classification rule-set and category definitions.

7. **Auditability**: Implement logging mechanisms to track the reasoning behind each categorization decision.

8. **Privacy-Preserving Classification**: Consider using homomorphic encryption or secure multi-party computation for privacy-sensitive features.

9. **Efficient Querying**: Design category to record mapping to allow for efficient querying of records based on category criteria.
