package main

import (
	"log"
	"os"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		log.Fatalf("failed to create chaincode: %s", err)
	}

	tls := shim.TLSProperties{
		Disabled: true,
	}

	server := &shim.ChaincodeServer{
		CCID:     os.Getenv("CORE_CHAINCODE_ID"),
		Address:  os.Getenv("CORE_CHAINCODE_ADDRESS"),
		CC:       chaincode,
		TLSProps: tls,
	}

	if err := server.Start(); err != nil {
		log.Fatalf("failed to start chaincode server: %s", err)
	}

}
