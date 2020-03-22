# go-iden3-crypto wasm
*Warning: this is an ongoing experimentation*

WASM wrappers for browser execution of [go-iden3-crypto](https://github.com/iden3/go-iden3-core).

## Wasm usage
To compile to wasm, inside the `wasm` directory, execute:
```
GOARCH=wasm GOOS=js go build -o go-iden3-crypto.wasm go-iden3-crypto-wasm-wrappers.go
```

Add the file `wasm_exec.js` in the directory:
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

To see the usage from javascript, check `index.js` file.

Run the http server that allows to load the `.wasm` file:
```
cd webtest && node server.js
```

## Benchmarks
Tested on a Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz, with 16GB of RAM

- In browser
```
WASM poseidon 1000 iterations took 2499 milliseconds.
WASM poseidon benchmark: 2.499ms

iden3js/poseidon 1000 iterations took 13075 milliseconds.
iden3js/poseidon benchmark: 13.075ms
```

- While in native Go
```
Go poseidon 1000 iterations took 142.220917ms.
Go poseidon benchmark: 142.22µs

BenchmarkPoseidon-4                9424        127337 ns/op
```

