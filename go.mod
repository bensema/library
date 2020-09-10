module github.com/bensema/library

go 1.14

require (
	github.com/facebook/ent v0.4.2
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/gogo/protobuf v1.3.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gomodule/redigo v1.8.2
	github.com/google/uuid v1.1.2 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin v0.11.0
	go.opentelemetry.io/contrib/instrumentation/net/http v0.11.0
	go.opentelemetry.io/otel v0.11.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.11.0
	go.opentelemetry.io/otel/sdk v0.11.0
	go.uber.org/zap v1.16.0
	golang.org/x/image v0.0.0-20200801110659-972c09e46d76 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
