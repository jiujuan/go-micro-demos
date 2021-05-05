## tracing 链路追踪

由于微服务个数多，调用复杂，这时候对微服务错误的诊断也变得复杂起来，怎么应对这种新情况？

为了应对这种新情况，面向微服务的诊断和分析系统就应运而生，包括集中式日志系统(Logging)，度量系统

(Metrics) 和分布式追踪系统（Tracing).

- Logging：用于记录日志信息。比如业务信息，错误信息。
- Metrics：记录聚合的信息。比如http的请求个数。
- Tracing：记录请求范围的信息，链路追中信息。比如一次 rpc 的调用过程和耗时。



三者的关系如下图：

![logging-metrics-tracing关系](../imgs/logging-metrics-tracing-03.png)

(From: https://peter.bourgon.org/blog/2017/02/21/metrics-tracing-and-logging.html)

上图所表示的每一个领域都有相应的软件来实现其功能：

1. Metrics - Prometheus

2. Logging - ELK

3. Tracing - jaeger, Zipkin，SkyWalking 等

   

traceing 相关的论文:

1. Dapper 分布式追中系统 [Dapper, a Large-Scale Distributed Systems Tracing Infrastructure](https://research.google/pubs/pub36356/) 

2. 采样分析 [Uncertainty in Aggregate Estimates from Sampled Distributed Traces](https://research.google/pubs/pub40378/)



由于市面上的追踪系统越来越多，为了解决分布式追踪系统 API 相互不兼容的问题，诞生了 [OpenTracing](https://opentracing.io/) 和 [opencensus](https://opencensus.io/)

规范，后面这 2 个规范又合并到了 CNCF 下的 [OpenTelemetry](https://opentelemetry.io/) 。

## jager 简介

[jager](https://github.com/jaegertracing) 是 Uber 开发的追踪系统，后来捐献给了 CNCF，所以它是符合 OpenTelemetry 规范的。



jager 的安装部署可以看这篇文章：[jager简单实践](https://www.cnblogs.com/jiujuan/p/13235748.html)。



下面我们就使用 uber 的 jaeger 这个链路追踪工具结合 go-micro 来开发。

## go-micro 例子

看看 go-micro 中的 opentracing 插件，它提供了 4 中 wrapper 方法分别用于不同的服务类型：

1. WrapHandler() server 中间件
2. WrapCall() call 中间件
3. WrapClient() client 中间件
4. WrapSubscriber()  订阅中间件



先写一个公共应用的包 tracer/jaeger.go 。



