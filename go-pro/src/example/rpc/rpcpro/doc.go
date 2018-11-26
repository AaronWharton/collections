// 模块化设计
// 将RPC分为三个部分：
// 1.服务端实现RPC方法：server.go
// 2.客户端调用RPC方法：client.go
// 3.制定RPC接口规范的api（最重要，解决耦合，统一设计）：api/api.go-->服务的名字、服务的方法列表、注册该服务的方法（register）
package main
