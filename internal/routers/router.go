package routers

import (
	"github.com/gin-gonic/gin"
	api2 "larmic/sonnen-charger-api/internal/routers/handlers"
)

func InitRouter(chargerIp string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", api2.GetOpenApi)

	api := r.Group("/api")
	api.GET("/settings", api2.GetSettings(chargerIp))
	api.GET("/", api2.GetOpenApi)

	return r
}
