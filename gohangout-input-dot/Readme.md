[https://github.com/childe/gohangout](https://github.com/childe/gohangout) 插件示例.

Golang 的 Plugin 文档参考 [https://tip.golang.org/pkg/plugin/](https://tip.golang.org/pkg/plugin/)

创建一个 input, 输入一个 `.` 字符

## 编译

将 dot.go 复制到 gohangout 主目录下面, 运行

```shell
go build -buildmode=plugin -o dot.so dot.go
```

## gohangout 配置示例

```yaml
inputs:
    - '/path/go/dot.so': {}
outputs:
    - Stdout: {}
```
