
1. 编译
```cmd
set GOOS=js
set GOARCH=wasm
go build -o yourgame.wasm

```

2. 复制js
```cmd
cp '%GOROOT%\misc\wasm\wasm_exec.js' .

```

3. html文档
```html
<!DOCTYPE html>
<script src="wasm_exec.js"></script>
<script>
// Polyfill
if (!WebAssembly.instantiateStreaming) {
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch("yourgame.wasm"), go.importObject).then(result => {
    go.run(result.instance);
});
</script>

```
4. 运行
```go
go run github.com/hajimehoshi/wasmserve@latest .
```