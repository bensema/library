package trace

import (
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/label"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
)

type Config struct {
	DSN         string
	ServiceName string
	Proto       string
}

// 初始化全局trace
func Init(c *Config) func() {
	hostname, _ := os.Hostname()
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(c.DSN),
		jaeger.WithProcess(jaeger.Process{
			ServiceName: c.ServiceName,
			Tags: []label.KeyValue{
				label.String("hostname", hostname),
			},
		}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)

	if err != nil {
		log.Fatal(err)
	}
	return flush
}
