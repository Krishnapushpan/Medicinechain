## Decentralized Medicine Selling System

project is a Decentralized Medicine Selling System built using Hyperledger Fabric. The main objective of this project is to ensure transparency, traceability, and trust in the pharmaceutical supply chain by using blockchain technology.

In this system, different participants such as manufacturers, dealer and shop interact on a secure and permissioned network. Each transaction—from manufacturing to selling—is recorded immutably on the blockchain, ensuring that medicines cannot be tampered with or counterfeited during the supply chain process.

This project includes smart contracts (chaincode) written in Go to manage medicine data.

## Initialize Go project
```bash
go mod init medicine
```
## Get the Fabric Contract API
```bash
go get github.com/hyperledger/fabric-contract-api-go@v1.2.1
```
## Get all required dependencies & remove unnecessary ones
```bash
go mod tidy
```
## Start the network using script file
```bash
./startNetwork.sh
```
## Prepare Minifab for deployment
```bash
sudo chmod -R 777 vars/
mkdir -p vars/chaincode/Medicine/go
cp -r ../Chaincode/* vars/chaincode/Medicine/go/
```
## Invoke chaincode
```bash
minifab invoke -n Medicine -p '"CreateMedicine","med01","Paracetamol","500mg","2025-12-31","Batch-01"'
```
## Query chaincode
```bash
minifab query -n Medicine -p '"ReadMedicine","med01"'
```

