package v1

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"testing"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func checkInit_v2(t *testing.T, stub *shim.MockStub, args []string) {
	//checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("123"), []byte("B"), []byte("234")})

	//res := stub.MockInit("1", args)
	var res peer.Response
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func conv1(arr [][]byte) []string {
	var args []string = make([]string, 0)
	//fmt.Printf("%v \n",args)
	for _, x := range arr {
		//fmt.Printf("%s \n",string(x))
		args = append(args, string(x))
	}
	return args

}

func conv2(x []string) ([][]byte){

	var y[][]byte = make([][]byte,0)
	for _,item := range x{
		by := []byte(item)
		//y[idx]=by
		y = append(y, by)
	}

	return y
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {

	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("query"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query", name, "failed to get value")
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func checkInvoke_v2(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", conv2(args))
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}


func Test_TokenInit(t *testing.T) {

	scc := new(TokenChainCode)

	stub := shim.NewMockStub("ex02", scc)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("123"), []byte("B"), []byte("234")})

	checkInit_v2(t, stub, []string{"init", "a", "123"})

	//checkState(t, stub, "A", "123")
	//checkState(t, stub, "B", "234")
}

func Test_TokenIco(t *testing.T) {
	var args []string = []string{}
	fmt.Printf("%v",args)
	scc := new(TokenChainCode)

	stub := shim.NewMockStub("ex02", scc)

	//checkInvoke(t, stub, [][]byte{[]byte("invoke"), []byte("A"), []byte("B"), []byte("123")})

	checkInvoke_v2(t,stub,[]string{"ico","token1","TKN1","1000","5"})


}

func TestToken_Query(t *testing.T) {
	scc := new(TokenChainCode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=345 B=456
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("345"), []byte("B"), []byte("456")})

	// Query A
	checkQuery(t, stub, "A", "345")

	// Query B
	checkQuery(t, stub, "B", "456")
}

func TestToken_Invoke(t *testing.T) {
	scc := new(TokenChainCode)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=567 B=678
	checkInit(t, stub, [][]byte{[]byte("init"), []byte("A"), []byte("567"), []byte("B"), []byte("678")})

	// Invoke A->B for 123
	checkInvoke(t, stub, [][]byte{[]byte("invoke"), []byte("A"), []byte("B"), []byte("123")})
	checkQuery(t, stub, "A", "444")
	checkQuery(t, stub, "B", "801")

	// Invoke B->A for 234
	checkInvoke(t, stub, [][]byte{[]byte("invoke"), []byte("B"), []byte("A"), []byte("234")})
	checkQuery(t, stub, "A", "678")
	checkQuery(t, stub, "B", "567")
	checkQuery(t, stub, "A", "678")
	checkQuery(t, stub, "B", "567")
}
