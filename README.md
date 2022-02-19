# On-boarding

基于`Golang 1.17.7`开发。

## 开发指南

```bash
# 安装swag生成器
$ go install github.com/swaggo/swag/cmd/swag@latest
# 初始化swag文档
$ swag init -d .,handler

# 安装air(可选, 用于热更新)
$ go get -u github.com/cosmtrek/air
# 初始化air(可选)
$ air init
# 使用air动项目
$ air

# 或者也可以直接使用go启动
$ go run .
```
