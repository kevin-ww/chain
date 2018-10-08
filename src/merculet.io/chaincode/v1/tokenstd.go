package v1

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

//https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md

type Account string

type TokenOwner struct {
	Name   string
	Symbol string
	Addr   Account
}

type TokenIssuer struct {
	Name string
	Addr Account
}

type Token struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    uint8  `json:"decimals"`
	TotalSupply uint64 `json:"total_supply"`
}

type TokenIssueStd interface {
	ico(name string, symbol string, decimals uint8, totalSupply uint64) bool
	//name() string
	//symbol() string
	//decimals() uint8
	//totalSupply() uint64
	getTokenIssued(symbol string) *Token
	getAllTokenIssued() []*Token
}

type TokenStd interface {
	balanceOf(owner Account) uint64
	transfer(to Account, tokens uint64) (success bool)
	transferFrom(from Account, to Account, tokens uint64) (success bool)
	approve(spender Account, tokens uint64) (success bool)
	allowance(owner Account, spender Account) (remaining uint64)
}

func buildTokenKey(symbol string) string {
	return strings.TrimSpace("token-issue-" + strings.ToLower(symbol))
}

func buildTokenOwnerKey(owner Account, symbol string) string {
	return strings.ToLower(strings.Join([]string{"tkn", string(owner), symbol}, "-"))

}

func (i *TokenIssuer) ico(name string, symbol string, decimals uint8, totalSupply uint64) bool {

	var tk string

	token := Token{
		Name:        name,
		Symbol:      symbol,
		Decimals:    decimals,
		TotalSupply: totalSupply,
	}

	//GetStub().PutState()
	fmt.Printf("%v is now issuing the token [%v]", i, token)
	bytes, _ := json.Marshal(token)
	tk = buildTokenKey(symbol)
	err := PutState(tk, bytes)
	if err != nil {
		log.Println(err.Error())
		return true
	}
	return false
}

func (i *TokenIssuer) getTokenIssued(symbol string) *Token {
	var tk string
	tk = buildTokenKey(symbol)
	bytes, e := GetState(tk)
	if e != nil {
		log.Println(e.Error())
		return nil
	}
	var token Token
	err := json.Unmarshal(bytes, &token)
	if err != nil {
		log.Println(e.Error())
		return nil
	}
	return &token
}

func (i *TokenIssuer) getAllTokenIssued() []*Token {
	log.Println("not implemented yet")
	return nil
}

func (t *TokenOwner) balanceOf(owner Account) uint64 {
	var tok string
	tok = buildTokenOwnerKey(owner, t.Symbol)
	bytes, e := GetStub().GetState(tok)
	if e != nil {
		log.Println("no such owner")
		return 0
	}
	//
	return binary.BigEndian.Uint64(bytes)
}

func (t *TokenOwner) transfer(to Account, tokens uint64) (success bool) {

	balance := t.getBalance()

	if balance < tokens {
		fmt.Printf("don't have enough tokens , %d < %d", balance, tokens)
	}

	//balance2 := getBalance(to, t.Symbol)


	//transaction start


	PutState(buildTokenOwnerKey(t.Addr,t.Symbol), nil)
	PutState(buildTokenOwnerKey(to, t.Symbol),nil)



	//
	return true
}

func (t *TokenOwner) getBalance() uint64 {
	return 0
}

func getBalance(account Account, symbol string) uint64 {
	return 0
}
