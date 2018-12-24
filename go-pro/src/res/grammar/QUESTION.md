# QUESTION

## 1.同个包下不同文件不能互相调用函数
环境： macOS + GoLand
描述：同一个包下（准确说是 main 包才会发生这种情况）不同文件函数引用正常，但在运行时报错：
  ```
  # command-line-arguments
  api/main.go:12:23: undefined: CreateUser
  api/main.go:14:34: undefined: Login
  ```
原因：编译问题。当点击 GoLand 运行按钮编译运行 main 函数的时候不会默认编译 main 函数引用的同一包下的函数。
解决：使用命令行编译运行当前引用的所有文件，运行程序，成功：
```
go run handlers.go main.go
```
