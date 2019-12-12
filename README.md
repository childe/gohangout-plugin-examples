[https://github.com/childe/gohangout](https://github.com/childe/gohangout) 插件示例.

Golang 的 Plugin 文档参考 [https://tip.golang.org/pkg/plugin/](https://tip.golang.org/pkg/plugin/)

每一个的使用参见各个目录下的Readme.md

目前的 Go 版本(1.13), 如果使用 Plugin, 需要保证 module 版本是一样的. 所以使用 Plugin 的时候需要用下面两种编译方法

- 将 Plugin 代码复制到 Gohangout 目录下使用 go build 编译 so 文件
- 将 Plugin Module 化, 单独编译. 然后 go get go get github.com/childe/gohangout@vx.x.x 下载同样版本的 gohangout 程序
