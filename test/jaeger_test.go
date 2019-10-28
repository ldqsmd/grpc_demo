package tests

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"log"
	"testing"
	"time"
)

func TestJaeger(t *testing.T) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: ":6831", // 替换host
		},
	}
	closer, err := cfg.InitGlobalTracer(
		"醒了",
	)
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	var ctx = context.TODO()
	span1, ctx := opentracing.StartSpanFromContext(ctx, "span_1")
	span1.Finish()

	time.Sleep(time.Second / 2)
	span11, _ := opentracing.StartSpanFromContext(ctx, "span_1-1")
	time.Sleep(time.Second / 2)
	span11.Finish()
	span12, _ := opentracing.StartSpanFromContext(ctx, "span_1-2")
	span12.Finish()

	defer closer.Close()
}
