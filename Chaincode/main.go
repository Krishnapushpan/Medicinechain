package main

import (
	"log"

	"medicinechain/contract" 

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	medicineContract := new(contracts.MedicineContract)

	chaincode, err := contractapi.NewChaincode(medicineContract)
	if err != nil {
		log.Panicf("Could not create medicine chaincode: %v", err)
	}

	err = chaincode.Start()
	if err != nil {
		log.Panicf("Failed to start medicine chaincode: %v", err)
	}
}
