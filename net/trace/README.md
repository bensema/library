##### 项目简介
使用 opentelemetry opentelemetry-go-contrib

用jaeger收集

gin，grpc 每一个对Inject和Extract，spanContext注入参数中，用于跨进程追踪

```
https://github.com/open-telemetry/opentelemetry-go-contrib
```
grpc注入
```
conn, err := grpc.Dial(":7777", grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(grpcotel.UnaryClientInterceptor(global.Tracer(""))),
		grpc.WithStreamInterceptor(grpcotel.StreamClientInterceptor(global.Tracer(""))),
	)
```

gin注入
```
initTracer()
	r := gin.New()
	r.Use(gintrace.Middleware("my-server"))
```
