package trace

import (
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"grpc_demo/util/conf"
	"io"
	"log"
)

var traceClose io.Closer

func JaegerConfigInit() {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           conf.Config.Trace.LogSpans,
			LocalAgentHostPort: conf.Config.Trace.LocalAgentHostPort, // 替换host
		},
	}
	c, err := cfg.InitGlobalTracer(
		conf.Config.Trace.ServerName,
	)
	traceClose = c
	if err != nil {
		log.Fatal("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	//defer closer.Close()

}

func TraceClose() {
	traceClose.Close()
}
