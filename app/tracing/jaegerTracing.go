package tracing

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
	"xyzApp/app/config"
)

func ConnectJaeger(cfg *config.AppConfig, log *logrus.Logger, serviceName string) (opentracing.Tracer, io.Closer) {
	config := jaegerConfig.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: fmt.Sprintf("%v:%v", cfg.Jaeger.Host, cfg.Jaeger.Port),
		},
	}

	tracer, closer, err := config.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		log.Fatalf("cant connect to jaeger : %v", err.Error())
	}

	log.Info("success connect to jaeger")
	return tracer, closer
}
