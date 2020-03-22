#!/bin/sh

GOARCH=wasm GOOS=js go build -o go-iden3-crypto.wasm go-iden3-crypto-wasm-wrappers.go && cp go-iden3-crypto.wasm webtest/ && mv go-iden3-crypto.wasm build/go-iden3-crypto.wasm
