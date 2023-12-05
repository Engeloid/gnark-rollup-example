# Rollup_TXs

Facilitation of rollup transactions with basic zk-rollups by gnark.

## Setup

### Run solc locally using docker to generate solidity contract

First bind local path, so that docker container has the repo to work with the files.

```bash
docker run -v .\contracts\:/sources ethereum/solc:stable -o /sources --abi --bin /sources/gnark_verifier.sol 
```

### Create Contract binding

Move the `abi` and `.bin` files to `bindings/` and then to create the contract binding:

```bash
abigen --abi .\bindings\Verifier.abi --pkg main --type Verifier --out .\src\Verifier.go
```

To then inject the contract binding `Verifier.go` for deployment:

```bash
abigen --abi .\bindings\Verifier.abi --pkg main --type Verifier --out .\src\Verifier.go --bin .\bindings\Verifier.bin 
```

## Truffle and Ganache for contract delpoyment

Deploy all contracts from scratch

```bash
truffle migrate --reset
```

## Run Go main and all other necessary go scripts

```bash
go run .\src\.
```
