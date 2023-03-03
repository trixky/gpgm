# Refs

- https://golangbot.com/webassembly-using-go/

# Compile to wasm

```bash
GOOS=js GOARCH=wasm go build -o ../client/static/wasm/src/main.wasm
```