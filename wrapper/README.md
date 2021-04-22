# wrapper: go-micro 的包装功能

Wrapper 提供了一种功能，可以在执行某方法前先执行 Wrapper 的代码，hook 思想还是 AOP 思想？
所以 wrapper 功能可以在客户端和服务器做很多功能：熔断限流、Filter、Auth等。可以看到在官方 go-plugins/wrapper

包里面，作者利用该机制封装了很多功能。



wrapper 怎么使用呢？ 看一下官方的代码 [logwrapper](https://github.com/micro/examples/examples/wrapper)。

