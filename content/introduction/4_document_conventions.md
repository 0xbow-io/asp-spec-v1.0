# 1.4 Document Conventions

## This document adheres to the following conventions:

### [1] Mathematical Expressions:

All mathematical formulas are rendered with [KaTeX](https://katex.org/)

### [2] Code Snippets:

Code examples are provided in syntax-highlighted blocks.

```solidity
pragma solidity ^0.8.0;

contract PublicRegistry {
    // Contract code here
}
```

### [2] Diagrams:

System diagrams and flowcharts are presented using either Mermaid.js or D2 syntax . For example:

```text
graph LR
    A[Off-Chain Components] --> B[On-Chain Components]
    B --> C[External Systems]
```

```mermaid

%%{
  init: {
    'theme': 'base',
    'themeVariables': {
      'primaryColor': '#1e1e2e',
      'primaryTextColor': '#cdd6f4',
      'primaryBorderColor': '#89b4fa',
      'lineColor': '#fab387',
      'secondaryColor': '#f9e2af',
      'tertiaryColor': '#a6e3a1',
      'noteTextColor': '#1e1e2e',
      'noteBkgColor': '#f5c2e7',
      'notesBorderColor': '#cba6f7',
      'textColor': '#cdd6f4',
      'fontSize': '16px',
      'labelTextColor': '#1e1e2e',
      'actorBorder': '#89b4fa',
      'actorBkg': '#1e1e2e',
      'actorTextColor': '#cdd6f4',
      'actorLineColor': '#89b4fa',
      'signalColor': '#cdd6f4',
      'signalTextColor': '#1e1e2e',
      'loopTextColor': '#cdd6f4',
      'activationBorderColor': '#f5c2e7',
      'activationBkgColor': '#1e1e2e',
      'sequenceNumberColor': '#1e1e2e'
    }
  }
}%%

graph LR
    A[Off-Chain Components] --> B[On-Chain Components]
    B --> C[External Systems]
```

### [3] Terminology:

Technical terms specific to the ASP system are defined in the glossary [Appendix 13.1](./appendices/1_glossary.md) and are italicized upon first use in each section.

### [4] References:

Citations to external documents or standards are provided in square brackets and listed in the References section [Appendix 13.2](./appendices/1_references.md).

### [5] Notes and Warnings:

Important information is highlighted in note blocks via [mdbook-admonish plugin](https://tommilligan.github.io/mdbook-admonish/overview.html).

```admonish note title='Note'
Note: Critical implementation details are emphasized in such blocks.
```

```admonish warning title='Warning'
Warning: Potential risks or security concerns are highlighted in warning blocks.
```

### [5] Version Information:

Any version-specific information is clearly marked with the applicable version number.
