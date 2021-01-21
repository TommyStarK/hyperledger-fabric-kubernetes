package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract ...
type SmartContract struct {
	contractapi.Contract
}

// Init ...
func (sc *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	var asset = &SimpleAsset{
		Content: "default",
		TxID:    ctx.GetStub().GetTxID(),
	}

	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %w", err)
	}

	return ctx.GetStub().PutState("default-asset", assetAsBytes)
}

// Delete ...
func (sc *SmartContract) Delete(ctx contractapi.TransactionContextInterface, key string) error {
	return ctx.GetStub().DelState(key)
}

// Query ...
func (sc *SmartContract) Query(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	bytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return "", fmt.Errorf("failed to query the ledger for specific key: %s", key)
	}

	// not found
	if bytes == nil || len(bytes) == 0 {
		return "", nil
	}

	return string(bytes), nil
}

// QueryPrivateData ...
func (sc *SmartContract) QueryPrivateData(ctx contractapi.TransactionContextInterface, key string) (string, error) {
	value, err := ctx.GetStub().GetPrivateData("dummy", key)
	if err != nil {
		return "", fmt.Errorf("Failed to retrieve value of the specified key '%s': %w", key, err)
	}

	// not found
	if value == nil || len(value) == 0 {
		return "", nil
	}

	return string(value), nil
}

// Store ...
func (sc *SmartContract) Store(ctx contractapi.TransactionContextInterface, key, stringifiedAsset string) error {
	var asset = &SimpleAsset{}
	if err := json.Unmarshal([]byte(stringifiedAsset), asset); err != nil {
		return err
	}

	asset.TxID = ctx.GetStub().GetTxID()
	assetAsBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset: %w", err)
	}

	return ctx.GetStub().PutState(key, assetAsBytes)
}

// StorePrivateData ...
func (sc *SmartContract) StorePrivateData(ctx contractapi.TransactionContextInterface) error {
	tmap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return err
	}

	for key, value := range tmap {
		if err := ctx.GetStub().PutPrivateData("dummy", key, value); err != nil {
			return fmt.Errorf("Failed to put key '%s' with value (%s) into the transaction's private writeset ", key, value)
		}
	}

	return nil
}

// SetEvent ...
func (sc *SmartContract) SetEvent(ctx contractapi.TransactionContextInterface, eventFilter, message string) error {
	return ctx.GetStub().SetEvent(eventFilter, []byte(message))
}
