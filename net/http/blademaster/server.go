package blademaster

import "github.com/gin-gonic/gin"

type ServerConfig struct {
}

func DefaultServer(conf *ServerConfig) *gin.Engine {
	engine := gin.Default()
	engine.Use(Trace())
	return engine
}
