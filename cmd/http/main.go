package main

import (
	"fmt"
	"io"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jcConfig "github.com/uber/jaeger-client-go/config"
	"todoapp/config"
	"todoapp/internal/app"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	tracer, closer := InitTracer("todoapp")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	svc := app.InitHTTPServer(cfg)

	svc.RegisterMiddleware()
	svc.RegisterHandler()
	svc.Start()
}

func InitTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jcConfig.Configuration{
		ServiceName: service,
		Sampler: &jcConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jcConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: os.Getenv("JAEGER_AGENT_HOST") + ":" + os.Getenv("JAEGER_AGENT_PORT"),
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("failed to init tracer, error: %s", err.Error()))
	}

	return tracer, closer
}
