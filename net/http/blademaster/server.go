package blademaster

import "github.com/gin-gonic/gin"

type ServerConfig struct {
}

func New() *gin.Engine {
	engine := gin.New()
	return engine
}

func Default() *gin.Engine {
	engine := gin.Default()
	return engine
}

func DefaultServer(conf *ServerConfig) *gin.Engine {
	engine := gin.New()
	engine.Use(Trace())
	return engine
}
