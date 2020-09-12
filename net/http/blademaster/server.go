package blademaster

import "github.com/gin-gonic/gin"

type ServerConfig struct {
	OpenTrace bool
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
	engine := gin.Default()
	if conf.OpenTrace {
		engine.Use(Trace())
	}
	return engine
}
