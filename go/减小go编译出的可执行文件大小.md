# 减小go编译出的可执行文件大小

- -s: 去掉符号信息。
- -w: 去掉DWARF调试信息。

```bash
go build -i -ldflags="-s -w"
```
