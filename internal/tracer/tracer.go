package tracer

import (
	cfg "github.com/ozoncp/ocp-suggestion-api/internal/config"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	jmetrics "github.com/uber/jaeger-lib/metrics"
)

func InitTracing() (io.Closer, error) {
	jcfg := jaegercfg.Configuration{
		ServiceName: "ocp-suggestion-api",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: cfg.Config.TracerAddr,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := jmetrics.NullFactory

	tracer, closer, err := jcfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)
	return closer, nil
}
