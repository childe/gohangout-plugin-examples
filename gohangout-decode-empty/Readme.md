[https://github.com/childe/gohangout](https://github.com/childe/gohangout) 插件示例.

Golang 的 Plugin 文档参考 [https://tip.golang.org/pkg/plugin/](https://tip.golang.org/pkg/plugin/)

创建一个 decoder, 将 input 进来的数据, 编译成一个空字典(所以这种东西实际情况上根本不会用到吧)

## 编译

将 empty_decoder.go 复制到 gohangout 主目录下面, 运行

```shell
go build -buildmode=plugin -o empty_decoder.so empty_decoder.go
```

## gohangout 配置示例

```yaml
inputs:
    - Stdin:
        codec: '/path/to/emptydecoder.so'
outputs:
    - Stdout: {}
```
