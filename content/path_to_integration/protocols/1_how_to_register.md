# Protocol / DApp Integration Pathway

```admonish info title='0xbow ASP V1.0 go-Buildkit'

0xbow ASP v1.0 `go-build-kit` is a set of primitive Go modules for building custom ASP solutions.

With the go-build-kit, you can easily:
  - Interact with exisiting ASP service
  - Integrate ASP modules into your protocol / DApp
  - Build & deploy your own custom ASP services

`go-build-kit` is open-source and will be available soon.
```

0xBow ASP v1.0 implements an extensible `Integration Framework` which offers a broad range of
functionality that can be readily integrated into custom solutions for your protocol / DApp.

0xBow ASP offers REST, gRPC and WebSocket [APIs](<(../../appendices/4_API.md)>)
to support offchain integration and onchain contracts for onchain integration.

If your requirements are not met by these existing APIs, you can register for a custom integration
with the ASP by following the steps below.

> **All integration efforts will contribute to the maturity & adoption of the ASP** ❤️

## How to Register?

> You can find all prior registrations in the [Protocol Registry](../../protocol_registry/) page.

A **Registration** is the aknowledgement of an integration request and marks
the beginning of the integration process. It is a formal step that allows both
parties (i.e 0xBow ASP and Protocol X) to track the progress of the
integration process.

To register, you will need to first to submit a new `Integration Request` issue in the
[asp-spec-v1.0a](https://github.com/0xbow-io/asp-spec-v1.0/issues) repository.

Be sure to specify the following details in your request:

- [x] **Integration Type**: `Protocol/dAPP Integration`
- [x] **Integration Target**: The name of your protocol / dApp (i.e. "Protocol X")
- [x] **Involvement**: What's your involvement with the protocol / dApp? (i.e. Developer / Engineer, Founder, etc.)
- [x] **Contact Information**: How can we reach you? (i.e. Email, Twitter, Telegram, etc.)
- [x] **Integration Description**: A brief description of your protocol / dApp and the integration requirements.

- ### What are the possible integration options?

  0xBow has taken a modular approach to the ASP implementation, allowing for external
  integrations to be made with ease.

  Your protocol / DApp can leverage independent ASP services & modules
  to suit your specific requirements, i.e.:

  - Utilize the ASP `Watcher` service to observe and record protocol / DApp state-transitions:
    - Integrate `observer` & `state-transition recorder` modules into your services.
    - Or subscribe to Watcher aWebSocket endpoints to receive event streams.
  - Utilize the ASP `Categorization Engine ` to categorize specific events:
    - Subscribe to Categorization Engine  WebSocket endpoints
    - Request the categorization of a `Record` via gRPC or REST API.
    - Integrate the categorization pipeline into your services.
  - Leverage the Onchain ASP `Public Registry` or Offchain `Record Archive` to support
    business rules or compliant privacy-preserving mechanics
    (i.e. public inputs to onchain verifier contracts).

  #### Here are some example use-cases:

  ***

  > **Use-case 1: Restricted ERC20 Token Airdrop**
  >
  > Protocol X is planning to `airdrop` ERC20 tokens to a restricted set of accounts.
  >
  > The conditions for the airdrop:
  >
  > - Account must have directly interacted with the Protocol.
  > - Account must have a minimum balance of 1 ETH.
  > - Account must have a minimum of 100 transactions.
  > - Account is not directly & indirectly associated with any illicit activities.

  **Integration Path:**

  - ASP generates the schema for `Airdrop Eligible` category with features reflecting the
    specified conditions.
  - ASP will record all protocol interactions, categorize them and publish the category
    bitmaps to the `Public Registry`.
  - ASP will deploy registry-adapter contract which contains a mapping of Account addresses &
    record hashes as well as the bitmap filter for `Airdrop Eligible`.
  - The Airdrop can now integrate with the registry-adapter to ensure that only eligible
    accounts receive the airdrop.

  ***

  > **Use-case 2: Compliant ERC-4337 Paymaster**
  >
  > Protocol Y wishes to implement compliant `ERC-4337 Paymaster`
  >
  > The compliance rules :
  >
  > - Account must have completed KYC verification.
  > - Account's UserOps are not associated with any illicit activities.

  **Integration Path:**:

  - ASP generates the schema for `COMPLIANT_ACCOUNT` category with features reflecting the
    specified conditions.
  - ASP will record all protocol interactions, categorize them and publish the category
    bitmaps to the `Public Registry`.
  - ASP will deploy registry-adapter contract which contains a mapping of Account addresses &
    record hashes as well as the bitmap filter for `COMPLIANT_ACCOUNT`.
  - The Paymster can now integrate the registry-adpater into it's `validatePaymasterUserOp` function
    to ensure that only compliant accounts can interact with the Paymaster.

  ***

- ### What details are required for the Integration Description?

  > Technical context on Protocol-ASP integration can be found in section
  > [3.1 Protocol Requirements](../protocol_integration/1_protocol_requirements.md).

  The `Integration Description` should provide a brief overview of your protocol / DApp
  and the integration requirements. This should include:

  - A brief description of your protocol / DApp
  - The integration requirements
  - Any specific features or functionalities that you would like to integrate
  - Any specific modules or services that you would like to leverage

  > The more detailed the `Integration Description`, the better we can understand your
  > requirements and provide a tailored integration solution.

- ### This is too confusing for me !

  If you are unsure about the registration process or how the ASP can be integrated
  with your protocol / DApp, feel free to reach out to us at [0xBow.io](https://0xbow.io/).

  We're happy to guide you through the process and answer any questions you may have.

- ### What happens after I submit the registration request?

  After submission, 0xBow will review the integration requirements, conduct workshop sessions
  to plan the integration process and deliver a detailed integration plan with timelines.
  Once complete, 0xBow will request for a signoff on the integration plan.

  Upon signoff, the integration request will be documented in the [Protocol Registry](../../protocol_registry/) page
  with links to the integration project tracking page.
