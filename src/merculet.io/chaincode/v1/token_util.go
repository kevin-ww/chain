package v1

import "github.com/hyperledger/fabric/core/chaincode/shim"

var instance *shim.MockStub

func GetStub() *shim.MockStub{
	if instance==nil{
		instance = shim.NewMockStub("tokenstd",nil)
	}
	return instance
}

func  PutState(key string, value []byte) error {
	stub := GetStub()
	stub.MockTransactionStart("tx01")
	err := stub.PutState(key, value)
	stub.MockTransactionEnd("tx01")
	return err
}

// GetState retrieves the value for a given key from the ledger
func  GetState(key string) ([]byte, error) {
	stub:=GetStub()
	return stub.GetState(key)
}




