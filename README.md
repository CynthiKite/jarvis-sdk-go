<div align="left">
<a href="https://indykite.com">
<img src="https://github.com/indykite/.github/blob/master/assets/IndyKITE_Rough_red.png" alt="IndyKite Red Logo" width="100px" height="183px" align="right">
</a>
</div>

# IndyKite Client Libraries for Go

IndyKite is a cloud identity platform built to secure and manage
human & non-person (IoT) identities and their data. This repository contains the
Golang Library packages for [IndyKite Platform](https://indykite.com) Client SDK.

[![Build](https://github.com/indykite/jarvis-sdk-go/actions/workflows/pr-test.yaml/badge.svg)](https://github.com/indykite/jarvis-sdk-go/actions/workflows/pr-test.yaml)
[![codecov](https://codecov.io/gh/indykite/jarvis-sdk-go/branch/master/graph/badge.svg?token=TFCDLXbnsh)](https://codecov.io/gh/indykite/jarvis-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/indykite/jarvis-sdk-go)](https://goreportcard.com/report/github.com/indykite/jarvis-sdk-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/indykite/jarvis-sdk-go.svg)](https://pkg.go.dev/github.com/indykite/jarvis-sdk-go)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Findykite%2Fjarvis-sdk-go.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Findykite%2Fjarvis-sdk-go?ref=badge_shield)

```go
import "github.com/indykite/jarvis-sdk-go"
```


## Used terminology

| Definition               | Description                                                                                      |
|--------------------------|--------------------------------------------------------------------------------------------------|
| Digital Twin             | A digital twin is the digital identity of a physical entity on/in a software/identity system     |
| Application Space ID     | ID of the application where the digital twin belongs to                                          |
| Application Agent ID     | ID of the agent which makes the application available for the different calls                    |
| Tenant ID                | ID of the tenant where the digital twin belongs to. The tenant is belong to an application space |
| Private Key and Settings | The secret which required to reach the system. Indykite provides the necessary secrets           |
| Property                 | The digital twin's property (eg.: email, name)                                                   |
| JWT                      | JSON Web Tokens                                                                                  |
| Introspect               | A process used to validate the token and to retrieve properties assigned to the token            |
| Patch property           | Add, change or delete a property of a digital twin                                               |

## Documentation

Visit the IndyKite One Developer Community site for official
[IndyKite documentation](https://indykite.one/blog?category=5e3e9297-3451-4b52-91ee-8027dcd1789c)
and to find out how to use the entire platform for your project.

## Getting Started

### Requirements

* Go 1.19
* AppAgent credentials and Service account credentials. These are obtained when setting up the Admin Console. 

### Install

```go
go get -u github.com/indykite/jarvis-sdk-go
```

### Configure the development environment
You need to have a configuration JSON file (AppAgent credentials or Service account credentials) to be able to use the Jarvis Proto SDK.
Example configuration file:

```json
{
  "appSpaceId": "696e6479-6b69-4465-8000-010100000002",
  "baseUrl": "https://jarvis.indykite.com",
  "applicationId": "696e6479-6b69-4465-8000-020100000002",
  "defaultTenantId": "696e6479-6b69-4465-8000-030100000002",
  "appAgentId": "696e6479-6b69-4465-8000-050100000002",
  "endpoint": "jarvis.indykite.com",
  "privateKeyJWK": {
    "kty": "EC",
    "d": "aa",
    "use": "sig",
    "crv": "P-256",
    "kid": "2e5lIxxb6obIwpok",
    "x": "6d83se2Eg",
    "y": "lshzMo",
    "alg": "ES256"
  },
  "privateKeyPKCS8Base64": "LS0tLS==",
  "privateKeyPKCS8": "-----BEGIN PRIVATE KEY-----\nM\n-----END PRIVATE KEY-----"
}
```

Conditionally optional parameters:

- baseUrl
- defaultTenantId
- endpoint

There are two ways to set up the necessary credentials. You either pass the JSON configuration file (AppAgent credentials or Service account credentials) to the `INDYKITE_APPLICATION_CREDENTIALS` environment variable or set the `INDYKITE_APPLICATION_CREDENTIALS_FILE` environment variable to the configuration file's path.

   1. on Linux and OSX

      ```shell
      export INDYKITE_APPLICATION_CREDENTIALS='{"appSpaceId":"00000000-0000-4000-a000-000000000000","appAgentId":"00000000-0000-4000-a000-000000000001","endpoint": "application.indykite.com","privateKeyJWK":{"kty":"EC","d": "abcdef","use": "sig","crv": "P-256","kid":"efghij","x":"klmnop","y":"qrstvw","alg":"ES256"}}'`
      ```

      or

      ```shell
      export INDYKITE_APPLICATION_CREDENTIALS_FILE=/Users/xx/configuration.json
      ```

   1. on Windows command line

      ```shell
      setex INDYKITE_APPLICATION_CREDENTIALS='{"appSpaceId":"00000000-0000-4000-a000-000000000000","appAgentId":"00000000-0000-4000-a000-000000000001","endpoint": "application.indykite.com","privateKeyJWK":{"kty":"EC","d": "abcdef","use": "sig","crv": "P-256","kid":"efghij","x":"klmnop","y":"qrstvw","alg":"ES256"}}'`
      ```

      or

      ```shell
      setex INDYKITE_APPLICATION_CREDENTIALS_FILE "C:\Users\xx\Documents\configuration.json"
      ```

## Support and Feedback

Feel free to file a bug, submit an issue or give us feedback on our
[issues page](https://github.com/indykite/jarvis-sdk-go/issues)

## Vulnerability Reporting

[Responsible Disclosure](responsible_disclosure.md)

## Changelog

[Changelog](CHANGELOG.md)

## Contributors / Acknowledgements

Coming soon!

## What is IndyKite

IndyKite is building the identity platform for Web 3.0.

IndyKite’s decentralized Identity Platform unlocks contextual insights from identity data to enable businesses to realize greater ROI.. With products that securely manage human, IoT,and machine identity, the IndyKite Identity Platform leverages machine learning to deliver context-aware authorization, knowledge driven decisions and risk analytics.

Built on a knowledge graph data model, IndyKite enables developers with flexible APIs through a growing open-source ecosystem.

## License

[This project is licensed under the terms of the Apache 2.0 license.](LICENSE)
