[https://github.com/childe/gohangout](https://github.com/childe/gohangout) 插件示例.

Golang 的 Plugin 文档参考 [https://tip.golang.org/pkg/plugin/](https://tip.golang.org/pkg/plugin/)

创建一个 output, 不管前面的数据是什么, 只是打印一个`-` (所以这种东西实际情况上根本不会用到吧)

## 编译

将 dash.go 复制到 gohangout 主目录下面, 运行

```shell
go build -buildmode=plugin -o dash.so dash.go
```

## gohangout 配置示例

```yaml
inputs:
    - Stdin: {}
outputs:
    - '/path/to/dash.so': {}
```
