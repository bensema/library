package blademaster

import (
	"github.com/gin-gonic/gin"
	gintrace "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin"
)

func Trace() gin.HandlerFunc {
	return gintrace.Middleware("")
	//return gintrace.Middleware("service")
}
