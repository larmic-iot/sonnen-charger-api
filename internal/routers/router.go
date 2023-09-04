package routers

import (
	"github.com/gin-gonic/gin"
	api2 "larmic/sonnen-charger-api/internal/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	api.GET("/settings", api2.GetSettings)

	return r
}
