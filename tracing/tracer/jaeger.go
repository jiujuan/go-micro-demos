package tracer

import (
    "github.com/opentracing/opentracing-go"
    "github.com/uber/jaeger-client-go"
    jaegercfg "github.com/uber/jaeger-client-go/config"
    "io"
    "time"
)

func NewJaeger(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
    config := jaegercfg.Configuration{
        ServiceName: serviceName,
        Sampler: &jaegercfg.SamplerConfig {
            Type: jaeger.SamplerTypeConst,
            Param: 1,
        },
        Reporter: &jaegercfg.ReporterConfig{
            LogSpans           : false,
            BufferFlushInterval: 1 * time.Second,
        },
    }

    sender, err := jaeger.NewUDPTransport(addr, 0)
    if err != nil {
        return nil, nil, err
    }

    reporter := jaeger.NewRemoteReporter(sender)
    tracer, closer, err := config.NewTracer(
        jaegercfg.Reporter(reporter),
    )

    return tracer, closer, err
}
