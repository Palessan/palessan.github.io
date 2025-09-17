# palessan.github.io

To generate wasm:
GOOS=js GOARCH=wasm go build -o main.wasm main.go

To run wasm:
wasmserve.exe .
http://localhost:8080/

If you get `illegal base64 data at input byte 3`
copy wasm_exec.js under same directory