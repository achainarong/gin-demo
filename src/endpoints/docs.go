package endpoints

import (
	"gin-demo/src/config"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// SetupSwagger sets up the Swagger endpoint
func SetupSwagger(r *gin.Engine, settings *config.Settings) {
	r.GET("/docs", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
