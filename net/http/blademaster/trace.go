package blademaster

import (
	"github.com/gin-gonic/gin"
)

func Trace() gin.HandlerFunc {
	return Middleware("")
	//return gintrace.Middleware(service, opts...)
}
