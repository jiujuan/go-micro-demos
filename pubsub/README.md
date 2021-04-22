# 发布订阅

## demo1下的例子

pub, sub 的例子，根据官方的例子修改了一下。

1.proto文件

message.proto

```proto
syntax = "proto3";

message Message {
    string id = 1;
    string message = 2;
}
```
运行命令生成对应 go 文件：
 
 `protoc --go_out=. --micro_out=. ./message.proto`

2.发布消息程序：
client/main.go

3.订阅消息程序：
server/main.go

4.运行 go run ./sever/main.go

5.运行 go run ./client/main.go

程序运行成功，会在终端打印发布、订阅的日志消息