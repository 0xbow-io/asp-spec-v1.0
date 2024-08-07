# 2.1 ASP Architecture

```d2
direction: right
style: {
  fill: transparent
}


title: Fig.2.1. High Level Architecture of the ASP System{
  near: bottom-center
  shape: text
  style: {
    font-size: 29
    bold: true
    underline: false
    font-color: white
  }
}


asp: {
  label: "ASP System"
  label.near: outside-top-left
  shape: rectangle
  style: {
    border-radius:10
    fill: transparent
    font-color: "#cdd6f4"
    bold: true
    font-size: 32
    stroke: "#cdd6f4"
  }

  ofc: {
    label: "Service Stack"
    label.near: outside-top-left
    shape: rectangle
    style: {
      border-radius:10
      fill: transparent
      stroke-dash: 3
      stroke-width: 2
      font-color: "#cdd6f4"
      stroke: "#cba6f7"
      bold: true
      font-size: 32
    }

    obs: {
      label: "Observer"
      label.near: outside-top-left
      shape: rectangle
      style: {
        border-radius:10
        fill: transparent
        font-color: "#cdd6f4"
        stroke: "#74c7ec"
        bold: true
        font-size: 32

      }
      Pm: {
        label: "watcher"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#b4befe"
          bold: true
          font-size: 32
        }
      }
      Std: {
        label: "State-Transition Detector"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#b4befe"
          bold: true
          font-size: 32
        }
      }
      Rg: {
        label: "Record-Generator"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#b4befe"
          bold: true
          font-size: 32
        }
      }

      Pm -> Std -> Rg {
        style: {
          stroke: "transparent"
          stroke-width: 0
          }
      }
    }
    Cg: {
      label: "Category-Engine"
      shape: rectangle
      label.near: outside-top-left
      style: {
        border-radius:10
        fill: transparent
        font-color: "#cdd6f4"
        stroke: "#74c7ec"
        bold: true
        font-size: 32
      }
      Fe: {
        label: "Feature-Extractor"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#b4befe"
          bold: true
          font-size: 32
        }
      }
      Cl: {
        label: "Classifier"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#b4befe"
          bold: true
          font-size: 32
        }
      }
      Cr: {
        label: "Categorizer"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#b4befe"
          bold: true
          font-size: 32
        }
      }

      Fe -> Cl -> Cr {
        style: {
          stroke: "transparent"
          stroke-width: 0
          }
      }

    }
  }


  onc: {
      label: "On-Chain Instances"
      label.near: outside-top-left
      shape: rectangle
      style: {
        border-radius:10
        fill: transparent
        font-color: "#cdd6f4"
        stroke: "#a6adc8"
        stroke-dash: 3
        stroke-width: 2
        bold: true
        font-size: 32
      }

      Prg: {
        label: "Public Registry"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#fab387"
          bold: true
          font-size: 32
        }
      }

      Zkp: {
        label: "ZKP Verifier"
        shape: rectangle
        style: {
          border-radius:10
          fill: transparent
          font-color: "#cdd6f4"
          stroke: "#fab387"
          bold: true
          font-size: 32
        }
      }
    }

    ofc -> onc {
      style: {
        stroke: "transparent"
        stroke-width: 0
        }
    }

}

```

> **The core idea of Privacy Pools is this:**
>
> Instead of merely proving that their withdrawal is
> linked to some previously-made deposit, a user proves membership in a more restrictive association set.
> This association set could be the full subset of previously-made deposits, a set consisting only of the user's
> own deposits, or anything in between ...
>
> Users will subscribe to intermediaries, called `association set providers (ASPs)`,
> which generate association sets that have certain properties" [^note]

The 0xBow ASP system is an implementation of the ASP concept initially introduced for Privacy Pool but now exteded to
facilitate compliance mechanisms across multiple blockchain protocols.

It's architecture is relatively simple and consists of two main components:
- `Service Stack`: 2 modular services that are working in concert to monitor, classify, and verify state transitions.
- `On-Chain Instances`: Components supporting onchain integrations with the ASP.

[^note]:
    ["Blockchain privacy and regulatory compliance: Towards a practical equilibrium"](https://www.sciencedirect.com/science/article/pii/S2096720923000519),
    Vitalik. B, Soleimani. A, et al., 2023
