package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var (
	testcc   *SmartContract
	teststub *shimtest.MockStub
)

func setup() {
	testcc, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		panic("failed to create chaincode: " + err.Error())
	}

	teststub = shimtest.NewMockStub("SmartContract", testcc)
}

func TestInitChaincode(t *testing.T) {
	result := teststub.MockInit("", nil)
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestInitLedger(t *testing.T) {
	result := teststub.MockInvoke("init-ledger", [][]byte{[]byte("Init")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryLedgerAfterInit(t *testing.T) {
	result := teststub.MockInvoke("query-ledger-after-init", [][]byte{[]byte("Query"), []byte("default-asset")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}

	var asset = &SimpleAsset{}
	if err := json.Unmarshal(result.Payload, asset); err != nil {
		t.Fatal(err)
	}

	if asset.Content != "default" {
		t.Error(`asset content should be a string with value "default"`)
	}

	if asset.TxID != "init-ledger" {
		t.Error(`asset transaction ID should be a string with value "init-ledger"`)
	}
}

func TestStoreAsset1(t *testing.T) {
	result := teststub.MockInvoke("store-asset1", [][]byte{[]byte("Store"), []byte("asset1"), []byte(`{"content": "foo"}`)})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryAsset1(t *testing.T) {
	result := teststub.MockInvoke("query-ledger-asset1", [][]byte{[]byte("Query"), []byte("asset1")})
	if result.Status != shim.OK {
		t.Errorf("failed to query the ledger: %s", result.GetMessage())
	}

	var asset = &SimpleAsset{}
	if err := json.Unmarshal(result.Payload, asset); err != nil {
		t.Fatal(err)
	}

	if asset.Content != "foo" {
		t.Error(`asset content should be a string with value "foo"`)
	}

	if asset.TxID != "store-asset1" {
		t.Error(`asset transaction ID should be a string with value "store-asset1"`)
	}
}

func TestUpdateAsset1(t *testing.T) {
	result := teststub.MockInvoke("update-asset1", [][]byte{[]byte("Store"), []byte("asset1"), []byte(`{"content": "bar", "txID": "store-asset1"}`)})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryUpdatedAsset1(t *testing.T) {
	result := teststub.MockInvoke("query-updated-asset1", [][]byte{[]byte("Query"), []byte("asset1")})
	if result.Status != shim.OK {
		t.Errorf("failed to query the ledger: %s", result.GetMessage())
	}

	var asset = &SimpleAsset{}
	if err := json.Unmarshal(result.Payload, asset); err != nil {
		t.Fatal(err)
	}

	if asset.Content != "bar" {
		t.Error(`asset content should be a string with value "bar"`)
	}

	if asset.TxID != "update-asset1" {
		t.Error(`asset transaction ID should be a string with value "update-asset1"`)
	}
}

func TestDeleteDefaultAsset(t *testing.T) {
	result := teststub.MockInvoke("delete-default-asset", [][]byte{[]byte("Delete"), []byte("default-asset")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}

func TestQueryUnknownAsset(t *testing.T) {
	result := teststub.MockInvoke("query-unknown-asset", [][]byte{[]byte("Query"), []byte("unknown-asset")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}

	if len(result.Payload) > 0 {
		t.Error("payload response should be empty")
	}
}

func TestSetEvent(t *testing.T) {
	result := teststub.MockInvoke("set-event", [][]byte{[]byte("SetEvent"), []byte("dummy"), []byte("test")})
	if result.Status != shim.OK {
		t.Fatal(result.GetMessage())
	}
}
func TestPrivateData(t *testing.T) {
	teststub.TransientMap = make(map[string][]byte)
	teststub.TransientMap["test-pdc"] = []byte("this is a test")

	resultStore := teststub.MockInvoke("test-store-private-data", [][]byte{[]byte("StorePrivateData")})
	if resultStore.Status != shim.OK {
		t.Error("chaincode invoke 'StorePrivateData' should have succeed")
	}

	resultQuery := teststub.MockInvoke("test-query-private-data", [][]byte{[]byte("QueryPrivateData"), []byte("test-pdc")})
	if resultQuery.Status != shim.OK {
		t.Error("chaincode invoke 'QueryPrivateData' should have succeed")
	}

	if string(resultQuery.Payload) != "this is a test" {
		t.Error(`payload as string should equal "this is a test"`)
	}
}

func TestFailureCases(t *testing.T) {
	result := teststub.MockInvoke("store-dummy", [][]byte{[]byte("Store"), []byte("dummy"), nil})
	if result.Status != shim.ERROR {
		t.Fatal(result.GetMessage())
	}

	resultQuery := teststub.MockInvoke("test-query-private-data-no-args", [][]byte{[]byte("QueryPrivateData")})
	if resultQuery.Status != shim.ERROR {
		t.Error("should have failed as we provided no args")
	}

	resultQuery = teststub.MockInvoke("test-query-private-data-unknown-key", [][]byte{[]byte("QueryPrivateData"), []byte("unknown-key")})
	if len(resultQuery.Payload) > 0 {
		t.Error("payload should be empty as we queried an unknown key")
	}
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
