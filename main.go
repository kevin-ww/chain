package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)


type Tokenx struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    uint8  `json:"decimals"`
	TotalSupply uint64 `json:"total_supply"`
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}


func StrToBytes(s string) []byte{
	return []byte(s)
}

func BytesToStr(arr []byte ) string{
	return string(arr[:])
}

func StrToUint64(s string) (uint64 ,error){
	i,err := strconv.ParseUint(s,10,64)
	return i,err
}

func StrToUint8(s string) (uint8,error){
	i,err:=strconv.ParseUint(s,10,8)
	return uint8(i),err
}

func main(){

	i, err := StrToUint8("380")
		//strconv.ParseInt("192", 10, 8)

	if err !=nil{
		fmt.Printf("%v", err)
	}

	fmt.Printf("%d",i)
}

func test(){

	//fmt.Printf("hahaha")
	var s string
	s="123"
	bytes := StrToBytes(s)
	fmt.Printf("%v \n",bytes)
	i,_ := StrToUint64(s)
	fmt.Printf("%s \n", i )
	//
	//u := StrToUint8(s)
	//fmt.Printf("%v \n",u)

}

