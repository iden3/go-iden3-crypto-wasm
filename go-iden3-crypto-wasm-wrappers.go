package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"syscall/js"

	"github.com/iden3/go-iden3-crypto/poseidon"
	"github.com/iden3/go-iden3-crypto/utils"
)

func main() {
	c := make(chan struct{}, 0)
	println("WASM go-iden3-crypto-wasm initialized")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("poseidonHash", js.FuncOf(poseidonHash))
	js.Global().Set("hash", js.FuncOf(hash))
	js.Global().Set("hashBytes", js.FuncOf(hashBytes))
}

func stringsToBigInts(s []string) []*big.Int {
	var r []*big.Int
	for i := 0; i < len(s); i++ {
		r = append(r, utils.NewIntFromString(s[i]))
	}
	return r
}

func poseidonHash(this js.Value, values []js.Value) interface{} {
	var inpStr []string
	err := json.Unmarshal([]byte(values[0].String()), &inpStr)
	if err != nil {
		return js.ValueOf(err.Error())
	}
	inp := stringsToBigInts(inpStr)

	if len(inp) != poseidon.T {
		return js.ValueOf(fmt.Errorf("input length should be %v", poseidon.T))
	}
	var inpPoseidon [poseidon.T]*big.Int
	copy(inpPoseidon[:], inp[:poseidon.T])

	r, err := poseidon.PoseidonHash(inpPoseidon)
	if err != nil {
		return js.ValueOf(err.Error())
	}

	return js.ValueOf(string(r.String()))
}

func hash(this js.Value, values []js.Value) interface{} {
	var inpStr []string
	err := json.Unmarshal([]byte(values[0].String()), &inpStr)
	if err != nil {
		return js.ValueOf(err.Error())
	}
	inp := stringsToBigInts(inpStr)

	r, err := poseidon.Hash(inp)
	if err != nil {
		return js.ValueOf(err.Error())
	}

	return js.ValueOf(string(r.String()))
}

func hashBytes(this js.Value, values []js.Value) interface{} {
	inp := []byte(values[0].String())
	r, err := poseidon.HashBytes(inp)
	if err != nil {
		return js.ValueOf(err.Error())
	}
	return js.ValueOf(string(r.String()))
}
