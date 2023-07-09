package ginfx

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Moudule = fx.Provide(provideGinEngine)

func provideGinEngine() *gin.Engine {
	return gin.Default()
}
