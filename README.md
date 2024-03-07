# Rollup_TXs

Facilitation of rollup transactions with basic zk-rollups by gnark.

## Setup

## Contracts

Ensure that the contracts are saved as utf8.

### Run solc locally using docker to generate solidity contract

First bind local path, so that docker container has the repo to work with the files.

```bash
docker run -v .\:/sources ethereum/solc:stable -o /sources/bindings --abi --bin /sources/contracts/Rollup.sol --overwrite 
```

### Create Contract binding

Move the `abi` and `.bin` files to `bindings/` and then to create the contract binding:

```bash
abigen --abi .\bindings\Rollup.abi --pkg main --type Rollup --out .\src\Rollup.go
```

To then inject the contract binding `Verifier.go` for deployment:

```bash
abigen --abi .\bindings\Rollup.abi --pkg main --type Rollup --out .\src\Rollup.go --bin .\bindings\Rollup.bin
```

#### When having depencies, use `--combined-json`

```bash
docker run -v ${PWD}:/sources ethereum/solc:stable --combined-json abi,bin,bin-runtime /sources/contracts/Rollup.sol -o /sources/bindings/RollupCombined.json
```

```bash
abigen --abi .\bindings\Rollup.abi --pkg main --type Rollup --out .\src\Rollup.go
```

```bash
abigen --combined-json ./bindings/RollupCombined.json/combined.json --pkg main --type Rollup --out ./src/Rollup.go
```

#### On laptop abigen:

```bash
C:\Users\docla\go\pkg\mod\github.com\ethereum\go-ethereum@v1.11.6\cmd\abigen
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

### Sol file updates
`.sol` files should be UTF-8, not UTF-8 with BOM!

The automaticly created `gnark_verifier.sol` comes with some issues:

- the `error` keyword was introduced in solidity v0.8.4 therefore `pragma solidity ^0.8.4;` is needed
- memory-safe keyword does not exist for the `assembly` instructions, i.e., it needs to be removed
