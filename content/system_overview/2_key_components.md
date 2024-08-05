# 2.2 Key Components

## 2.2.1 Service Stack:

### 2.2.1.i. Observer:

```admonish info
Visit the [protocol integration](/content/protocol_integration/1_protocol_requirements.md) section for more details on
Protocol $\leftrightarrows$ ASP integration.
```

```d2
direction: right
style: {
  fill: transparent
}

title: Fig.2.2. observation flow{
  near: bottom-center
  shape: text
  style: {
    font-size: 64
    bold: true
    underline: false
    font-color: white
  }
}


Protocol: {
    label: "Protocol"
    label.near: outside-top-left
    shape: square
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#cdd6f4"
      bold: true
      font-size: 64
    }

    PreState:  {
      label: "s"
      shape: circle
      height: 400
      style: {
        border-radius:10
        fill: "#1e1e2e"
        font-color: "#cdd6f4"
        stroke: "#cdd6f4"
        bold: true
        font-size: 64
      }
    }
    PostState:  {
      label: "s'"
      shape: circle
      height: 400
      style: {
        border-radius:10
        fill: "#1e1e2e"
        font-color: "#cdd6f4"
        stroke: "#cdd6f4"
        bold: true
        font-size: 64
      }
    }
    PreState -> PostState  {
    label: "T(s,i)→s’"
      style: {
        stroke: "#f5e0dc"
        stroke-dash: 3
        stroke-width: 10
        bold: true
        font-size: 64
        font-color:"#cdd6f4"
        animated: true
      }
    }
}
obs: {
  label: "Observation Pipeline"
  label.near: outside-top-left
  shape: queue
  height: 1000
  style: {
    border-radius:10
    fill: transparent
    font-color: "#cdd6f4"
    stroke: "#74c7ec"
    bold: true
    font-size: 64
  }
  Pm: {
    label: "watcher"
    shape: rectangle
    height: 400
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#b4befe"
      bold: true
      font-size: 64
    }
  }

  Std: {
    label: "State-Transition Detector"
    label.near: outside-top-left
    shape: rectangle
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#b4befe"
      bold: true
      font-size: 64
    }
    f: {
      label: "ΔS(s,s’)→δ"
      shape: rectangle
      height: 400
      style: {
        border-radius:10
        fill: transparent
        font-color: "#cdd6f4"
        stroke: "#89b4fa"
        stroke-dash: 3
        stroke-width: 2
        bold: true
        font-size: 64
      }
    }
  }

  Sbfr: {
    label: "State Buffer"
    shape: cylinder
    width: 400
    height: 400
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#b4befe"
      font-size: 64
    }
  }

  Rg: {
    label: "Record-Generator"
    label.near: outside-top-left
    shape: rectangle
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#b4befe"
      bold: true
      font-size: 64
    }

    h: {
      label: "H(s,e)→h"
      shape: rectangle
      height: 400
      style: {
        border-radius:10
        fill: transparent
        font-color: "#cdd6f4"
        stroke: "#89b4fa"
        stroke-dash: 3
        stroke-width: 2
        bold: true
        font-size: 64
      }
    }
  }



  Pm -> Std: "[e, s']"{
    style: {
      stroke: "#f5e0dc"
      stroke-dash: 3
      font-color: "#cdd6f4"
      stroke-width: 10
      animated: true
      bold: true
      font-size: 64
    }
  }

  Std -> Sbfr: "[e,s,s']"{
    style: {
      stroke: "#f5e0dc"
      stroke-dash: 3
      font-color: "#cdd6f4"
      stroke-width: 10
      animated: true
      font-size: 64
      bold: true
    }
  }

  Std <- Sbfr: "s"{
    style: {
      stroke: "#f5e0dc"
      stroke-dash: 3
      font-color: "#cdd6f4"
      stroke-width: 10
      animated: true
      font-size: 64
      bold: true
    }
  }

  Sbfr -> Rg: "[e,s,s']"{
    style: {
      stroke: "#f5e0dc"
      stroke-dash: 3
      font-color: "#cdd6f4"
      stroke-width: 10
      animated: true
      bold: true
      font-size: 64
    }
  }
}

Protocol -> obs: "s'" {
  style: {
    stroke: "#f5e0dc"
    stroke-dash: 3
    font-color: "#cdd6f4"
    stroke-width: 10
    animated: true
    font-size: 64
    bold: true
  }
}

```

The Observer is a service that monitors & records the state-changes of specific protocols in real-time.
It is comprised of the following modules:

1. **Watcher**: Watches the network for signals (i.e. event emissions) by protocols that indicate a state-change has occured.s

   It requires protocol-specific components ( `adapter` & a `parser` ) in order to interface with the protocol
   at the network level.

   It is possible to have a 1 to 1 implementation of the `watcher` module for each protocol.
   or a single `watcher` module that can be configured to monitor multiple protocols via pluggable adapters & parsers.

2. **State Transition Detector**: Identifies and validates the state transitions signalled by the `watcher` module.

   The Detector is notified of the state-change by the `watcher` module.
   In response, it attemps to rebuild a representation of the state
   from data cached in the state buffer.

   > The `state buffer` is a ring-buffer for efficient caching of new states and quick retrieval of old states.

   And then compares the current state with the previous state usung the protocol's state comparator function:
   $$\Delta S(s, s') \rightarrow \delta$$

   Where $\delta \not = 0$ indicates a state transition.

3. **Record Generator**: Creates cryptographic records of state transitions, defined as the tuple:

   $$R = (Scope, e, h, h')$$

   Where:

   - $Scope$ is the unique identifier of the protocol instance
   - $e$ is the reference to the state transition event (i.e. the block number, transaction hash)
   - $h$ is the pre-state hash
   - $h'$ is the post-state hash
   - Where $h$ & $h'$ is computed by a state-hash function: $H(s, e) \rightarrow h$
   - $h \not = h'$

### 2.2.1.ii. Category Engine:

```admonish info
The `Categorization Engine` is crucial to the ASP's ability in supporting compliance mechanisms & allow end-users to generate
"Associaiton Sets".

The objective of the categorization is to correctly identify attributes or properties (expressed as categories)
of the state-transition event which are relevant to the compliance requirements specified by the protocol (or other entities)

Visit the [classification & categorization](/content/classification_and_categorization/1_rule_based_classification.md)
and [feature extraction](/content/feature_extraction/1_feature_extraction.md) sections for more details on the categorization process.
```

The Categorization Engine executes a FIFO pipeline of feature-extraction, classification & categorization algorithms
to categorise the state-transition event referenced by the record $R$.

The output is a 256-bit vector termed "category bitmap" or $B$, where each bit represents a specific category.

```d2
direction: right
style: {
  fill: transparent
}

title: Fig.2.3. Categorization Flow{
  near: bottom-center
  shape: text
  style: {
    font-size: 64
    bold: true
    underline: false
    font-color: white
  }
}


obs: {
  label: "Observer"
  shape: circle
  style: {
    border-radius:10
    fill: transparent
    font-color: "#cdd6f4"
    stroke: "#74c7ec"
    bold: true
    font-size: 64
  }
}

cat: {
  label: "Categorization Pipeline"
  label.near: outside-top-left
  shape: queue
  height: 600
  style: {
    border-radius:10
    fill: transparent
    font-color: "#cdd6f4"
    stroke: "#74c7ec"
    bold: true
    font-size: 64
  }
  fe: {
    label: "Feature-Extractors"
    shape: step
    height: 300
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#74c7ec"
      bold: true
      font-size: 64
      multiple: true
    }
  }

  cl: {
    label: "classifiers"
    shape: step
    height: 300
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#74c7ec"
      bold: true
      font-size: 64
      multiple: true
    }
  }

  cg: {
    label: "categorizers"
    shape: step
    height: 300
    style: {
      border-radius:10
      fill: transparent
      font-color: "#cdd6f4"
      stroke: "#74c7ec"
      bold: true
      font-size: 64
      multiple: true
    }
  }


  cl -> cg:  "[categories]"{
    style: {
      stroke: "#f5e0dc"
      stroke-dash: 3
      font-color: "#cdd6f4"
      stroke-width: 10
      animated: true
      font-size: 64
      bold: true
    }
  }

  fe -> cl: "[features]"{
    style: {
      stroke: "#f5e0dc"
      stroke-dash: 3
      font-color: "#cdd6f4"
      stroke-width: 10
      animated: true
      bold: true
      font-size: 64
    }
  }
}



rq: {
  label: "R Queue"
  shape: queue
  style: {
    border-radius:10
    fill: transparent
    font-color: "#cdd6f4"
    stroke: "#74c7ec"
    bold: true
    font-size: 64
  }
}

obs -> rq -> cat: "R"{
  style: {
    stroke: "#f5e0dc"
    stroke-dash: 3
    font-color: "#cdd6f4"
    stroke-width: 10
    animated: true
    bold: true
    font-size: 64
  }
}



bit: {
  label: "B queue"
  shape: queue
  style: {
    border-radius:10
    fill: transparent
    font-color: "#cdd6f4"
    stroke: "#74c7ec"
    bold: true
    font-size: 64
    multiple: true
  }
}

cat -> bit: "B"  {
  style: {
    stroke: "#f5e0dc"
    stroke-dash: 3
    font-color: "#cdd6f4"
    stroke-width: 10
    animated: true
    bold: true
    font-size: 64
  }
}

```

The Category Pipeline is the sequential execution of the following routines:

1. **Feature Extractors**: Routines that Analyzes records to extract relevant features for classification.

2. **Classifiers**: Routines which generates categories for records through rule-based computation over extracted feature sets.

3. **Categorizers**: Routines that geenerates a 256-bit category bitmap for each classified record.

### 2.2.1.iii. On-Chain Instances:

```admonish info
Visit the [public registry](/content/public_registry/1_smart_contract_specification.md) section for more details.
```

7. **Public Registry**: On-chain smart contract storing record hashes and category bitmaps.

The on-chain public registry supports on-chain intergration with the ASP, allowing for the following:

- Generation of Association Sets whilst preserving end-user privacy.
- Direct integration with protocol contracts for compliance verification.
- Serve public input directly to onchain verifies during a transaction that requires compliance verification.

8. **ZKP Verifier**: OnChain component that verifies zero-knowledge proofs of compliance.

   See [Zero-Knowledge Proofs](/content/zero_knowledge_proofs/1_zkp_system_overview.md) section for more details
