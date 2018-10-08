package v1

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

type TokenChainCode struct {
}

//type Account string

//type Token struct {
//	Name        string `json:"name"`
//	Symbol      string `json:"symbol"`
//	TotalSupply uint64 `json:"total_supply"`
//	Decimals    uint8  `json:"decimals"`
//}

type tokenStd interface {
	totalSupply() uint64
	balanceOf(owner Account) uint64
	allowance(owner Account, spender Account) (remaining uint64)
	transfer(owner Account, tokens uint64) (success bool)
	approve(spender Account, tokens uint64) (success bool)
	transferFrom(from Account, to Account) (success bool)
}

const (
	DEFAULT_DECIMALS uint8  = 18
	MAX_TOTAL_SUPPLY uint64 = 1<<64 - 1
)

func (s *TokenChainCode) Name() string {
	return "TOKEN"
}

func (s *TokenChainCode) Init(api shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Init the chaincode: " + s.Name())
	return shim.Success(nil)
}

func (s *TokenChainCode) Invoke(api shim.ChaincodeStubInterface) peer.Response {
	function, args := api.GetFunctionAndParameters()
	switch function {
	case "ico":
		return s.ico(api, args)
	case "totalSupply":
		return s.totalSupply(api, args)
	case "balanceOf":
		return s.balanceOf(api, args)
	case "allowance":
		return s.allowance(api, args)
	}
	return shim.Error("Invalid function name " + function)
}

func (s *TokenChainCode) ico(api shim.ChaincodeStubInterface, args []string) peer.Response {

	var token Token
	var name, symbol string
	var totalSupply uint64
	var decimals uint64

	name = args[0]
	symbol = args[1]
	totalSupply, _ = strconv.ParseUint(args[2], 10, 64)
	decimals, _ = strconv.ParseUint(args[3], 10, 8)

	token = Token{
		Name:        name,
		Symbol:      symbol,
		TotalSupply: totalSupply,
		Decimals:    uint8(decimals),
	}

	tokenAsBytes, _ := json.Marshal(token)

	e := api.PutState(s.Name()+symbol, tokenAsBytes)

	if e != nil {
		return shim.Error(e.Error())
	}

	return shim.Success([]byte("total supply:1000"))
}

func (s *TokenChainCode) totalSupply(api shim.ChaincodeStubInterface, args []string) peer.Response {

	var symbol string
	var tkn Token

	symbol=args[0]

	tokenAsBytes, e := api.GetState(s.Name() + symbol)

	if e!=nil{
		return shim.Error(e.Error())
	}

	json.Unmarshal(tokenAsBytes, &tkn)

	return shim.Success([]byte(string(tkn.TotalSupply)))
}

func (s *TokenChainCode) balanceOf(api shim.ChaincodeStubInterface, args []string) peer.Response {
	var addr ,symbol string

	var k string
	addr = args[0]
	symbol = args[1]

	k = addr+symbol

	balanceAsBytes, e := api.GetState(k)

	if e != nil{
		return shim.Error(e.Error())
	}

	balance ,_:= strconv.ParseUint(string(balanceAsBytes),10,64)

	return shim.Success([]byte(string(balance)))
}

func (s *TokenChainCode) allowance(api shim.ChaincodeStubInterface, args []string) peer.Response {
	return shim.Success([]byte("false"))
}



//func (s *TokenChainCode) set(api shim.ChaincodeStubInterface, key string, value []byte) error {
//	err := api.PutState(key, value)
//	if err != nil {
//		return fmt.Errorf("Failed to set asset: %s", key)
//	}
//	return nil
//}
//
//func (s *TokenChainCode) get(api shim.ChaincodeStubInterface, key string) ([]byte, error) {
//	value, err := api.GetState(key)
//	if err != nil {
//		return nil, fmt.Errorf("Failed to get asset: %s with error: %s", key, err)
//	}
//	if value == nil {
//		return nil, fmt.Errorf("Asset not found: %s", key)
//	}
//	return value, nil
//}

func (t *Token) transfer(owner Account, tokens uint64) (success bool){
	return true
}

func (t *Token) allowance(owner Account, spender Account) (remaining uint64){
	fmt.Println("not implemented")
	return 0
}

func (t *Token) approve(spender Account, tokens uint64) (success bool){
	return true
}

func (t *Token) transferFrom(from Account, to Account) (success bool){
	return true
}
