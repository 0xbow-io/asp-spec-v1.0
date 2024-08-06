# 1.1 Purpose

```admonish abstract title='ASP Specification V1.0 '

**This document serves to:**

1. Provide a comprehensive integration guide for `protocols`, `end-users` and `compliance systems.`
2. Establish a `standardized framework` for `building` and `operating` an ASP.
3. Define the `technical specifications` and `interfaces` for each component of the ASP architecture.
4. Outline `best practices` for ensuring `security`, `scalability`, and `privacy` in ASP implementations.

```

The Association-Set Provider (ASP) Specification V1.0 devloped by 0xBow.io defines a standardized
framework for implementing and operating or integrating with the ASP system

The ASP is designed to support compliance mechanisms for blockchain protocols, enabling
the verification of compliance with regulatory requirements and business rules.

The ASP system aims to enable `privacy-preserving compliance` for blockchain protocols such
as Privacy Pool, by leveraging zero-knowledge proofs (ZKPs) and
efficient data categorization techniques.

```admonish warning title='Current WIPs'

**This document is still a work in progress.**

Here is curent TODO list:

- [ ] Complete [7. Zero-Knowledge Proofs](../zero_knowledge_proofs/1_zkp_system_overview.md)
- [ ] Complete [8. End-User Integrations](./end_user_integration/1_entity_requirements.md)
- [ ] Complete [9. Compliance Policies](../compliance_policies/1_policy_definition_language.md)
- [ ] Complete [10 Scalability and Performance](../scalability_and_performance/1_sharding_strategies.md)
- [ ] Complete [11 Security and Auditing](../security_and_auditing/1_threat_model.md)
- [ ] Complete [12 Governance and Upgrades](../governance_and_upgrades/1_governance_model.md)


```

## Content Overview:

See below to find the section that best suits your needs.

- ### Association Set Provider (ASP) Architecture:

  - Begin freading from [2.1 ASP Architecture](../system_overview/1_asp_architecture.md) to
    understand the ASP architecture and its components.

- ### 0xBow ASP Developments:

  - 0xBow ASP [Roadmap](../0xbow-asp-v1.0/4_API.md).
  - View the [API](../0xbow-asp-v1.0/4_API.md) section for current progress of the 0xBow ASP API.
  - View the [Contracts](../0xbow-asp-v1.0/4_API.md) section for current progress of 0xBow ASP Contracts.

- ### Integrating a protocol with the ASP:

  - To get started with integrating a protocol with the ASP and vice versa, access the
    [Protocol Integration](../protocol_integration/1_protocol_requirements.md) section.

  - For technical context on protocol-level integrations, refer to
    [3.1 Protocol Requirements](../protocol_integration/1_protocol_requirements.md).

  - To view current integrations or integration-proposal submissions, refer to the
    [Protocol Registry](../protocol_registry) section.

- ### User onboarding:

  - If you are an end-user looking to utilise the ASP or you are onboarding end-users
    to a platform that utilises ASP services, refer to
    [8.2 Interacting with Protocols](end_user_integration/2_interacting_with_protocols.md)

- ### ASP to ASP Interoperability:

  - If you are operating/implementing an ASP designed with specifications different to 0xBow ASP V1.0 and
    wish to integrate with the 0xBow ASP system, please vierw the
    [ASP Interoperability](../path_to_integration/asp/README.md) section.

- ### Contributors:

  - To view the list of contributors to the ASP Specification, refer to the
    [Contributors](../misc/CONTRIBUTORS.md) section.

- ### Specification Version History:

  - To view the version history & changelog of the ASP Specification, refer to the
    [Version History](../misc/VERSION.md) section.
