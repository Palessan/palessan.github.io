# palessan.github.io

To generate:
GOOS=js GOARCH=wasm go build -o main.wasm main.go

To run:
/c/code/go/src/wasmserve/wasmserve.exe .
http://localhost:8080/

If you get `illegal base64 data at input byte 3`
copy wasm_exec.js under same directory