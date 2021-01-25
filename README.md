## Cassis
[![SecHack365](https://img.shields.io/badge/SecHack365-2020-ffd700.svg)](https://sechack365.nict.go.jp/)
### Cassis is CAsual Self Sovereign Identity System
Cassis is a system that provides a comprehensive infrastructure for DID.

The goal is to create an environment that is DID ready in a casual way.

### Usage
1. Clone this repository.
1. `go install`
1. Prepare a configuration file.
   - requirements : See [example](example/cassis_build_example)
      - working directory (i.e. `sandbox`)
      - file written in DOT language
1. `cassis build` — build Ledger and Agent from DOT language
   - Answer the questions that are prompted.
1. `cassis network invite` — Issue an invitation from the Issuer
   - Answer the questions that are prompted.
  
### System Requirements
- docker
- docker-compose

### Release
#### v0.1
- Added support for generating docker-compose.yml for Agent from DOT language
  - You can do the same or better than the [run_docker script](https://github.com/hyperledger/aries-cloudagent-python/blob/master/scripts/run_docker) in ACA-Py, but more easily.
- Notice : You will need to modify some of the source code to fit your environment :)
  - Set the third argument of `ConvertFromGraph` in [build.go](cmd/build.go) to your localhost IP address.