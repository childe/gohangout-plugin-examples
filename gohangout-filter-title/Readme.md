[https://github.com/childe/gohangout](https://github.com/childe/gohangout) 插件示例.

Golang 的 Plugin 文档参考 [https://tip.golang.org/pkg/plugin/](https://tip.golang.org/pkg/plugin/)

## 编译

将 title.go 复制到 gohangout 主目录下面, 运行

```shell
go build -buildmode=plugin -o title.so title.go
```

## gohangout 配置示例 

```yaml
inputs:
    - Stdin: {}

filters:
    - 'path/to/title.so':
        fields: [message]

outputs:
    - Stdout: {}
```

## 代码说明

### New

一定要有 New 函数, 定义如下: `New(config map[interface{}]interface{}) interface{}`

在 Gohangout 里面会调用 plugin 的 New 方法来生成一个 Filter 实例.

### Filter

Filter 定义如下 `Filter(event map[string]interface{}) (map[string]interface{}, bool)` , 实现 Gohangout 里面的 Filter 接口
